import {Result} from "./result";
import ArticleModel from "../model/article";

export interface ArticleResult extends Result<ArticleModel> {
}

export interface ArticleListResult extends Result<Array<ArticleModel>> {
}