export class ResetPassword {
    password: string;
    password2: string;
    requestId: string;
    constructor() {
        this.password = '';
        this.password2 = '';
        this.requestId = '';
    }
}