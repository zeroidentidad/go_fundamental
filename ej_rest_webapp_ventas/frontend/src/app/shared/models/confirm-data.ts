export class ConfirmData{
    title?:string;
    message?:string;
    showCancel?:boolean;

    constructor(title: string, message: string, showCancel: boolean){
        this.title=title;
        this.message=message;
        this.showCancel=showCancel;
    }
}