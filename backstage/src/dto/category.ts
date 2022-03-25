export class CategoryAddDTO {
    private _name: string;

    constructor(name: string = "") {
        this._name = name;
    }

    get name(): string {
        return this._name;
    }

    set name(value: string) {
        this._name = value;
    }
}

export class CategoryDeleteDTO {
    private _id: number;

    constructor(id: number = 0) {
        this._id = id;
    }


    get id(): number {
        return this._id;
    }

    set id(value: number) {
        this._id = value;
    }
}


