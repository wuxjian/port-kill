export namespace main {
	
	export class ProcessInfo {
	    protocol: string;
	    local_address: string;
	    remote_address: string;
	    state: string;
	    pid: string;
	
	    static createFrom(source: any = {}) {
	        return new ProcessInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.protocol = source["protocol"];
	        this.local_address = source["local_address"];
	        this.remote_address = source["remote_address"];
	        this.state = source["state"];
	        this.pid = source["pid"];
	    }
	}

}

