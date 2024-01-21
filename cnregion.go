package cnregion

import (
	_ "embed"
	"errors"
	"strings"
)

type CnRegion struct{}

func NewCnRegion() *CnRegion {
	return &CnRegion{}
}

func (r *CnRegion) ParseByCode(code string) (province, city, area, street string, err error) {
	if len(code) != 2 && len(code) != 4 && len(code) != 6 && len(code) != 9 {
		return "", "", "", "", errors.New("编码长度错误")
	}

	// 按照省、市、区、街道的顺序查找
	region, err := r.findRegionByCode(regions, code[:2])
	if err != nil {
		return "", "", "", "", err
	}
	province = region.Name
	if len(code) >= 4 {
		region, err = r.findRegionByCode(region.Children, code[:4])
		if err != nil {
			return "", "", "", "", err
		}
		city = region.Name
	}
	if len(code) >= 6 {
		region, err = r.findRegionByCode(region.Children, code[:6])
		if err != nil {
			return "", "", "", "", err
		}
		area = region.Name
	}
	if len(code) >= 9 {
		region, err = r.findRegionByCode(region.Children, code[:9])
		if err != nil {
			return "", "", "", "", err
		}
		street = region.Name
	}

	return province, city, area, street, nil
}

// ParseByName 根据省市区街道反向解析出编码
func (r *CnRegion) ParseByName(province, city, area string, street ...string) (code string, err error) {
	names := []string{province, city, area}
	names = append(names, street...)
	region, err := r.findRegionByName(regions, names, 0)

	return region.Code, err
}

// Search 通过关键字搜索省市区街道编码
func (r *CnRegion) Search(keyword string) (result []Region) {
	searchRegions(regions, keyword, "", &result)
	return result
}

// findRegionByCode 递归查找给定的代码
func (r *CnRegion) findRegionByCode(regions []Region, code string) (Region, error) {
	for _, region := range regions {
		if region.Code == code {
			return region, nil
		}
	}
	return Region{}, errors.New("给定的编码无效")
}

// findRegionByName 递归查找给定的名称
func (r *CnRegion) findRegionByName(regions []Region, names []string, level int) (Region, error) {
	if level >= len(names) {
		return Region{}, errors.New("省市区街道参数不足")
	}

	for _, region := range regions {
		if region.Name == names[level] {
			if level == len(names)-1 {
				return region, nil
			}
			return r.findRegionByName(region.Children, names, level+1)
		}
	}
	return Region{}, errors.New("给定的省市区街道无效")
}

// searchRegions 递归搜索函数
func searchRegions(regions []Region, keyword, prefix string, result *[]Region) {
	for _, region := range regions {
		current := prefix + region.Name
		if strings.Contains(current, keyword) {
			*result = append(*result, region)
		}
		// 继续搜索子地区
		searchRegions(region.Children, keyword, current, result)
	}
}
