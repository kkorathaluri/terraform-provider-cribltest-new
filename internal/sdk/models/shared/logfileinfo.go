// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type LogFileInfo struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

func (o *LogFileInfo) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *LogFileInfo) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}
