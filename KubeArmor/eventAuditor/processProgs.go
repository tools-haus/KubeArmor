// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Authors of KubeArmor

package eventauditor

import lbpf "github.com/kubearmor/libbpf"

// KubeArmor Event Auditor Programs
const (
	KAEASysExecveProg      KABPFProgName    = "ka_ea_sched_process_exec"
	KAEASysExecveEvent     KABPFEventName   = "sched:sched_process_exec"
	KAEASysExecveTailProg0 KABPFProgName    = "ka_ea_sched_process_exec_0"
	KAEASysExecveTailProg1 KABPFProgName    = "ka_ea_sched_process_exec_1"
	KAEASysExecveProgFile  KABPFObjFileName = "ka_ea_process.bpf.o"

	KAEASysExitProg     KABPFProgName    = "ka_ea_sched_process_exit"
	KAEASysExitEvent    KABPFEventName   = "sched:sched_process_exit"
	KAEASysExitProgFile KABPFObjFileName = "ka_ea_process.bpf.o"
)

// KAEAGetProg Function
func KAEAGetProg(name KABPFProgName) KABPFProg {
	switch name {
	case KAEASysExecveProg:
		return KABPFProg{
			Name:      KAEASysExecveProg,
			EventName: KAEASysExecveEvent,
			EventType: lbpf.KABPFLinkTypeTracepoint,
			TailProgs: []KABPFTailProg{
				{
					Name:  KAEASysExecveTailProg0,
					Index: 0,
				},
				{
					Name:  KAEASysExecveTailProg1,
					Index: 1,
				},
			},
			FileName: KAEASysExecveProgFile,
		}
	case KAEASysExitProg:
		return KABPFProg{
			Name:      KAEASysExitProg,
			EventName: KAEASysExitEvent,
			EventType: lbpf.KABPFLinkTypeTracepoint,
			TailProgs: []KABPFTailProg{},
			FileName:  KAEASysExitProgFile,
		}
	default:
		return KABPFProg{
			Name:      "",
			EventName: "",
			EventType: lbpf.KABPFLinkTypeUnspec,
			TailProgs: []KABPFTailProg{},
			FileName:  "",
		}
	}
}