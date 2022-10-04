package photon

import (
	"reflect"

	"github.com/google/uuid"
)

type IPCHubCallback func(ipcID string)

type IPCHubEventsMap map[string][]IPCHubCallback

type IPCMap map[string]*IPC

type _IPCHub struct {
	OnEvents   IPCHubEventsMap
	OnceEvents IPCHubEventsMap
	IPCS       IPCMap
}

func (ipcHub *_IPCHub) AddIPC(ipc *IPC) {
	id := uuid.New()
	idString := id.String()
	ipcHub.IPCS[idString] = ipc
	ipc.ID = idString

	ipcHub.TriggerCallbacks("add", ipc.ID)
}

func (ipcHub *_IPCHub) RemoveIPC(ipc *IPC) {
	idString := ipc.ID
	delete(ipcHub.IPCS, idString)

	ipcHub.TriggerCallbacks("remove", ipc.ID)
}

func (ipcHub *_IPCHub) TriggerCallbacks(event string, ipcID string) {
	onEvents := ipcHub.OnEvents
	onceEvents := ipcHub.OnceEvents

	for _, v := range onEvents[event] {
		v(ipcID)
	}

	for _, v := range onceEvents[event] {
		v(ipcID)
	}

	delete(onceEvents, event)
}

func (ipcHub *_IPCHub) On(event string, callback IPCHubCallback) func() {
	return ipcHub.RegisterEvent(event, callback, "on")
}

func (ipcHub *_IPCHub) Once(event string, callback IPCHubCallback) func() {
	return ipcHub.RegisterEvent(event, callback, "once")
}

func (ipcHub *_IPCHub) RegisterEvent(event string, callback IPCHubCallback, t string) func() {
	var eventsMap IPCHubEventsMap
	if t == "on" {
		eventsMap = ipcHub.OnEvents
	} else {
		eventsMap = ipcHub.OnceEvents
	}

	if _, ok := eventsMap[event]; !ok {
		var eventArray []IPCHubCallback
		eventsMap[event] = eventArray
	}

	eventsMap[event] = append(eventsMap[event], callback)

	return func() {
		for i, v := range eventsMap[event] {
			p1 := reflect.ValueOf(v).Pointer()
			p2 := reflect.ValueOf(callback).Pointer()
			if p1 == p2 {
				eventsMap[event] = remove(eventsMap[event], i)
				if len(eventsMap[event]) == 0 {
					delete(eventsMap, event)
				}

				return
			}
		}
	}
}

func (ipcHub *_IPCHub) GetIPC(ipcID string) *IPC {
	return ipcHub.IPCS[ipcID]
}
