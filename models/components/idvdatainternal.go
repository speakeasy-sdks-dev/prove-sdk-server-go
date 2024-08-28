// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

type IDVDataInternal struct {
	DataSource1        *DataSourceInternal `json:"dataSource1,omitempty"`
	DataSource2        *DataSourceInternal `json:"dataSource2,omitempty"`
	MultiCIPConfidence *string             `json:"multiCIPConfidence,omitempty"`
}

func (o *IDVDataInternal) GetDataSource1() *DataSourceInternal {
	if o == nil {
		return nil
	}
	return o.DataSource1
}

func (o *IDVDataInternal) GetDataSource2() *DataSourceInternal {
	if o == nil {
		return nil
	}
	return o.DataSource2
}

func (o *IDVDataInternal) GetMultiCIPConfidence() *string {
	if o == nil {
		return nil
	}
	return o.MultiCIPConfidence
}
