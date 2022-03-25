class TagModel {
    private _id: number = 0;
    private _name: string = "";


    get id(): number {
        return this._id;
    }

    get name(): string {
        return this._name;
    }
}

export default TagModel;