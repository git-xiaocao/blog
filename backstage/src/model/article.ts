import Base from "./base";
import Category from "./category";
import Tag from "./tag";

class ArticleModel extends Base {
    private _categoryId: number = 0;
    private _category: Category = new Category();
    private _title: string = "";
    private _content: string = "";
    private _pathMark: string = "";
    private _tags: Array<Tag> = new Array<Tag>();

    get categoryId(): number {
        return this._categoryId;
    }

    get category(): Category {
        return this._category;
    }

    get title(): string {
        return this._title;
    }

    get content(): string {
        return this._content;
    }

    get pathMark(): string {
        return this._pathMark;
    }

    get tags(): Array<Tag> {
        return this._tags;
    }
}

export default ArticleModel;