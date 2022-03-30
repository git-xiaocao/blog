import {Result} from "./result";

export interface PageResult<T> extends Result<PageData<T>> {
}

interface PageData<T> {
    pageCurrent: number,
    pageTotal: number,
    total: number,
    list: Array<T>,
}