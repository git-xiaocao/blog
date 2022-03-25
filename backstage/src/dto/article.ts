export class ArticleAddDTO {
    private _categoryId: number;
    private _title: string;
    private _content: string;
    private _pathMark: string;
    private _tagIds: Array<number>;
    private _newTags: Array<string>;


    constructor(
        categoryId: number = 0,
        title: string = "",
        content: string = "",
        pathMark: string = "",
        tagIds: Array<number> = new Array<number>(),
        newTags: Array<string> = new Array<string>(),
    ) {
        this._categoryId = categoryId;
        this._title = title;
        this._content = content;
        this._pathMark = pathMark;
        this._tagIds = tagIds;
        this._newTags = newTags;
    }


    get categoryId(): number {
        return this._categoryId;
    }

    set categoryId(value: number) {
        this._categoryId = value;
    }

    get title(): string {
        return this._title;
    }

    set title(value: string) {
        this._title = value;
    }

    get content(): string {
        return this._content;
    }

    set content(value: string) {
        this._content = value;
    }

    get pathMark(): string {
        return this._pathMark;
    }

    set pathMark(value: string) {
        this._pathMark = value;
    }

    get tagIds(): Array<number> {
        return this._tagIds;
    }

    set tagIds(value: Array<number>) {
        this._tagIds = value;
    }

    get newTags(): Array<string> {
        return this._newTags;
    }

    set newTags(value: Array<string>) {
        this._newTags = value;
    }
}

export class ArticleUpdateDTO extends ArticleAddDTO {
    private _id: number;

    constructor(
        id: number = 0,
        categoryId: number = 0,
        title: string = "",
        content: string = "",
        pathMark: string = "",
        tagIds: Array<number> = new Array<number>(),
        newTags: Array<string> = new Array<string>(),
    ) {
        super(categoryId, title, content, pathMark, tagIds, newTags);
        this._id = id;
    }

    get id(): number {
        return this._id;
    }

    set id(value: number) {
        this._id = value;
    }
}

export class ArticleDeleteDTO {
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



