package main

type ServerStatus struct {
	Asserts struct {
		Msg       float64 `json:"msg,omitempty"`
		Regular   float64 `json:"regular,omitempty"`
		Rollovers float64 `json:"rollovers,omitempty"`
		User      float64 `json:"user,omitempty"`
		Warning   float64 `json:"warning,omitempty"`
	} `json:"asserts,omitempty"`
	BackgroundFlushing struct {
		AverageMs    float64 `json:"average_ms,omitempty"`
		Flushes      float64 `json:"flushes,omitempty"`
		LastFinished string  `json:"last_finished,omitempty"`
		LastMs       float64 `json:"last_ms,omitempty"`
		TotalMs      float64 `json:"total_ms,omitempty"`
	} `json:"backgroundFlushing,omitempty"`
	Connections struct {
		Available    float64 `json:"available,omitempty"`
		Current      float64 `json:"current,omitempty"`
		TotalCreated float64 `json:"totalCreated,omitempty"`
	} `json:"connections,omitempty"`
	Cursors struct {
		ClientCursorsSize float64 `json:"clientCursors_size,omitempty"`
		Note              string  `json:"note,omitempty"`
		Pinned            float64 `json:"pinned,omitempty"`
		TimedOut          float64 `json:"timedOut,omitempty"`
		TotalNoTimeout    float64 `json:"totalNoTimeout,omitempty"`
		TotalOpen         float64 `json:"totalOpen,omitempty"`
	} `json:"cursors,omitempty"`
	Dur struct {
		Commits            float64 `json:"commits,omitempty"`
		CommitsInWriteLock float64 `json:"commitsInWriteLock,omitempty"`
		Compression        float64 `json:"compression,omitempty"`
		EarlyCommits       float64 `json:"earlyCommits,omitempty"`
		JournaledMB        float64 `json:"journaledMB,omitempty"`
		TimeMs             struct {
			Dt               float64 `json:"dt,omitempty"`
			PrepLogBuffer    float64 `json:"prepLogBuffer,omitempty"`
			RemapPrivateView float64 `json:"remapPrivateView,omitempty"`
			WriteToDataFiles float64 `json:"writeToDataFiles,omitempty"`
			WriteToJournal   float64 `json:"writeToJournal,omitempty"`
		} `json:"timeMs,omitempty"`
		WriteToDataFilesMB float64 `json:"writeToDataFilesMB,omitempty"`
	} `json:"dur,omitempty"`
	ExtraInfo struct {
		HeapUsageBytes float64 `json:"heap_usage_bytes,omitempty"`
		Note           string  `json:"note,omitempty"`
		PageFaults     float64 `json:"page_faults,omitempty"`
	} `json:"extra_info,omitempty"`
	GlobalLock struct {
		ActiveClients struct {
			Readers float64 `json:"readers,omitempty"`
			Total   float64 `json:"total,omitempty"`
			Writers float64 `json:"writers,omitempty"`
		} `json:"activeClients,omitempty"`
		CurrentQueue struct {
			Readers float64 `json:"readers,omitempty"`
			Total   float64 `json:"total,omitempty"`
			Writers float64 `json:"writers,omitempty"`
		} `json:"currentQueue,omitempty"`
		LockTime  float64 `json:"lockTime,omitempty"`
		TotalTime float64 `json:"totalTime,omitempty"`
	} `json:"globalLock,omitempty"`
	Host          string `json:"host,omitempty"`
	IndexCounters struct {
		Accesses  float64 `json:"accesses,omitempty"`
		Hits      float64 `json:"hits,omitempty"`
		MissRatio float64 `json:"missRatio,omitempty"`
		Misses    float64 `json:"misses,omitempty"`
		Resets    float64 `json:"resets,omitempty"`
	} `json:"indexCounters,omitempty"`
	LocalTime string `json:"localTime,omitempty"`
	Locks     map[string]struct {
		TimeAcquiringMicros struct {
			RGlobal float64 `json:"R,omitempty"`
			WGlobal float64 `json:"W,omitempty"`
			R       float64 `json:"r,omitempty"`
			W       float64 `json:"w,omitempty"`
		} `json:"timeAcquiringMicros,omitempty"`
		TimeLockedMicros struct {
			RGlobal float64 `json:"R,omitempty"`
			WGlobal float64 `json:"W,omitempty"`
			R       float64 `json:"r,omitempty"`
			W       float64 `json:"w,omitempty"`
		} `json:"timeLockedMicros,omitempty"`
	} `json:"locks,omitempty"`
	Mem struct {
		Bits              float64 `json:"bits,omitempty"`
		Mapped            float64 `json:"mapped,omitempty"`
		MappedWithJournal float64 `json:"mappedWithJournal,omitempty"`
		Resident          float64 `json:"resident,omitempty"`
		Supported         bool    `json:"supported,omitempty"`
		Virtual           float64 `json:"virtual,omitempty"`
	} `json:"mem,omitempty"`
	Metrics struct {
		Cursor struct {
			Open struct {
				NoTimeout float64 `json:"noTimeout,omitempty"`
				Pinned    float64 `json:"pinned,omitempty"`
				Total     float64 `json:"total,omitempty"`
			} `json:"open,omitempty"`
			TimedOut float64 `json:"timedOut,omitempty"`
		} `json:"cursor,omitempty"`
		Document struct {
			Deleted  float64 `json:"deleted,omitempty"`
			Inserted float64 `json:"inserted,omitempty"`
			Returned float64 `json:"returned,omitempty"`
			Updated  float64 `json:"updated,omitempty"`
		} `json:"document,omitempty"`
		GetLastError struct {
			Wtime struct {
				Num         float64 `json:"num,omitempty"`
				TotalMillis float64 `json:"totalMillis,omitempty"`
			} `json:"wtime,omitempty"`
			Wtimeouts float64 `json:"wtimeouts,omitempty"`
		} `json:"getLastError,omitempty"`
		Operation struct {
			Fastmod      float64 `json:"fastmod,omitempty"`
			Idhack       float64 `json:"idhack,omitempty"`
			ScanAndOrder float64 `json:"scanAndOrder,omitempty"`
		} `json:"operation,omitempty"`
		QueryExecutor struct {
			Scanned        float64 `json:"scanned,omitempty"`
			ScannedObjects float64 `json:"scannedObjects,omitempty"`
		} `json:"queryExecutor,omitempty"`
		Record struct {
			Moves float64 `json:"moves,omitempty"`
		} `json:"record,omitempty"`
		Repl struct {
			Apply struct {
				Batches struct {
					Num         float64 `json:"num,omitempty"`
					TotalMillis float64 `json:"totalMillis,omitempty"`
				} `json:"batches,omitempty"`
				Ops float64 `json:"ops,omitempty"`
			} `json:"apply,omitempty"`
			Buffer struct {
				Count        float64 `json:"count,omitempty"`
				MaxSizeBytes float64 `json:"maxSizeBytes,omitempty"`
				SizeBytes    float64 `json:"sizeBytes,omitempty"`
			} `json:"buffer,omitempty"`
			Network struct {
				Bytes    float64 `json:"bytes,omitempty"`
				Getmores struct {
					Num         float64 `json:"num,omitempty"`
					TotalMillis float64 `json:"totalMillis,omitempty"`
				} `json:"getmores,omitempty"`
				Ops            float64 `json:"ops,omitempty"`
				ReadersCreated float64 `json:"readersCreated,omitempty"`
			} `json:"network,omitempty"`
			Preload struct {
				Docs struct {
					Num         float64 `json:"num,omitempty"`
					TotalMillis float64 `json:"totalMillis,omitempty"`
				} `json:"docs,omitempty"`
				Indexes struct {
					Num         float64 `json:"num,omitempty"`
					TotalMillis float64 `json:"totalMillis,omitempty"`
				} `json:"indexes,omitempty"`
			} `json:"preload,omitempty"`
		} `json:"repl,omitempty"`
		Storage struct {
			Freelist struct {
				Search struct {
					BucketExhausted float64 `json:"bucketExhausted,omitempty"`
					Requests        float64 `json:"requests,omitempty"`
					Scanned         float64 `json:"scanned,omitempty"`
				} `json:"search,omitempty"`
			} `json:"freelist,omitempty"`
		} `json:"storage,omitempty"`
		Ttl struct {
			DeletedDocuments float64 `json:"deletedDocuments,omitempty"`
			Passes           float64 `json:"passes,omitempty"`
		} `json:"ttl,omitempty"`
	} `json:"metrics,omitempty"`
	Network struct {
		BytesIn     float64 `json:"bytesIn,omitempty"`
		BytesOut    float64 `json:"bytesOut,omitempty"`
		NumRequests float64 `json:"numRequests,omitempty"`
	} `json:"network,omitempty"`
	Ok         float64 `json:"ok,omitempty"`
	Opcounters struct {
		Command float64 `json:"command,omitempty"`
		Delete  float64 `json:"delete,omitempty"`
		Getmore float64 `json:"getmore,omitempty"`
		Insert  float64 `json:"insert,omitempty"`
		Query   float64 `json:"query,omitempty"`
		Update  float64 `json:"update,omitempty"`
	} `json:"opcounters,omitempty"`
	OpcountersRepl struct {
		Command float64 `json:"command,omitempty"`
		Delete  float64 `json:"delete,omitempty"`
		Getmore float64 `json:"getmore,omitempty"`
		Insert  float64 `json:"insert,omitempty"`
		Query   float64 `json:"query,omitempty"`
		Update  float64 `json:"update,omitempty"`
	} `json:"opcountersRepl,omitempty"`
	Pid         float64                `json:"pid,omitempty"`
	Process     string                 `json:"process,omitempty"`
	RecordStats map[string]interface{} `json:"recordStats,omitempty"`
	Repl        struct {
		Arbiters   []string `json:"arbiters,omitempty"`
		Hosts      []string `json:"hosts,omitempty"`
		Ismaster   bool     `json:"ismaster,omitempty"`
		Me         string   `json:"me,omitempty"`
		Primary    string   `json:"primary,omitempty"`
		Secondary  bool     `json:"secondary,omitempty"`
		SetName    string   `json:"setName,omitempty"`
		SetVersion float64  `json:"setVersion,omitempty"`
	} `json:"repl,omitempty"`
	Uptime           float64 `json:"uptime,omitempty"`
	UptimeEstimate   float64 `json:"uptimeEstimate,omitempty"`
	UptimeMillis     float64 `json:"uptimeMillis,omitempty"`
	Version          string  `json:"version,omitempty"`
	WriteBacksQueued bool    `json:"writeBacksQueued,omitempty"`
}
type RecordStats struct {
	AccessesNotInMemory       float64 `json:"accessesNotInMemory,omitempty"`
	PageFaultExceptionsThrown float64 `json:"pageFaultExceptionsThrown,omitempty"`
}
