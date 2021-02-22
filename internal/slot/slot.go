package slot

// MaxSlotSize 最大槽数量
const MaxSlotSize = 2<<13 - 1

// Slot 槽
type Slot uint16

// RangeSlot 槽范围
type RangeSlot struct {
	MinSlot Slot `json:"minSlot"`
	MaxSlot Slot `json:"maxSlot"`
}
