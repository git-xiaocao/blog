//用户注册
export class UserRegisterDTO {
    private _email: string;
    private _password: string;
    private _code: string;


    constructor(
        email: string = "",
        password: string = "",
        code: string = "",
    ) {
        this._email = email;
        this._password = password;
        this._code = code;
    }


    get email(): string {
        return this._email;
    }

    set email(value: string) {
        this._email = value;
    }

    get password(): string {
        return this._password;
    }

    set password(value: string) {
        this._password = value;
    }

    get code(): string {
        return this._code;
    }

    set code(value: string) {
        this._code = value;
    }
}

//用户修改密码
export class UserChangePasswordDTO {
    private _email: string;
    private _oldPassword: string;
    private _newPassword: string;

    constructor(
        email: string = "",
        oldPassword: string = "",
        newPassword: string = "",
    ) {
        this._email = email;
        this._oldPassword = oldPassword;
        this._newPassword = newPassword;
    }


    get email(): string {
        return this._email;
    }

    set email(value: string) {
        this._email = value;
    }

    get oldPassword(): string {
        return this._oldPassword;
    }

    set oldPassword(value: string) {
        this._oldPassword = value;
    }

    get newPassword(): string {
        return this._newPassword;
    }

    set newPassword(value: string) {
        this._newPassword = value;
    }
}

//用户找回密码
export class UserRetrievePasswordDTO {
    private _email: string;
    private _code: string;
    private _newPassword: string;

    constructor(
        email: string = "",
        code: string = "",
        newPassword: string = "",
    ) {
        this._email = email;
        this._code = code;
        this._newPassword = newPassword;
    }


    get email(): string {
        return this._email;
    }

    set email(value: string) {
        this._email = value;
    }

    get code(): string {
        return this._code;
    }

    set code(value: string) {
        this._code = value;
    }

    get newPassword(): string {
        return this._newPassword;
    }

    set newPassword(value: string) {
        this._newPassword = value;
    }
}


