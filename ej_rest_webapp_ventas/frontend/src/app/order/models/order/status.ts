export class Status {
    StatusId: number;
    StatusName: string;

    static CreateInstance(statusId: number, statusName: string): Status {
        let status: Status=new Status();
        status.StatusId=statusId;
        status.StatusName=statusName;
        return status;
    }
}