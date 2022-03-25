import Base from "./base";

class UserModel extends Base {
    private _nick: string = "";
    private _email: string = "";
    private _password: string = "";

    get nick(): string {
        return this._nick;
    }

    get email(): string {
        return this._email;
    }

    get password(): string {
        return this._password;
    }
}

export default UserModel;