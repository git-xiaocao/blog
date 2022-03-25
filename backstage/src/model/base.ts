class Base {
    private _id: number = 0;
    private _createAt: string | null = null;
    private _updateAt: string | null = null;


    get id(): number {
        return this._id;
    }

    get createAt(): string | null {
        return this._createAt;
    }

    get updateAt(): string | null {
        return this._updateAt;
    }
}

export default Base;