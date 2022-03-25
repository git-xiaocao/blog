import {Result} from "./result";
import CategoryModel from "../model/category";

export interface CategoryListResult extends Result<Array<CategoryModel>> {
}