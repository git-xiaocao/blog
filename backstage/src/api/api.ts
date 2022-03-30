import {ArticleAddDTO, ArticleUpdateDTO, ArticleDeleteDTO} from "../dto/article";
import {EmptyResult} from "../vo/empty";
import {ArticleListResult, ArticleResult} from "../vo/article";
import {CategoryAddDTO, CategoryDeleteDTO} from "../dto/category";
import {TagAddDTO, TagDeleteDTO} from "../dto/tag";
import {CategoryListResult} from "../vo/category";
import {TagListResult} from "../vo/tag";
import axios, {AxiosInstance} from "axios";

namespace API {
    const axiosInstance: AxiosInstance = axios.create({
        baseURL: process.env.REACT_APP_BASE_URL,
    });

    export class Article {
        public static Add = (dto: ArticleAddDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/article/add", dto).then(res => res.data);
        }

        public static Update = (dto: ArticleUpdateDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/article/update", dto).then(res => res.data);
        }

        public static Delete = (dto: ArticleDeleteDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/article/delete", dto).then(res => res.data);
        }

        public static GetByPathMark = (pathMark: string): Promise<ArticleResult> => {
            return axiosInstance.get<ArticleResult>(`/article/${pathMark}`).then(res => res.data);
        }

        /**
         * 获取文章分页
         * @param pageCurrent 当前页
         * @param pageSize 每页显示数量最大30
         * @param tagIds 标签id数组
         * */
        public static GetPage = (pageCurrent: number, pageSize: number, tagIds: Array<number> | null): Promise<ArticleListResult> => {
            return axiosInstance.get<ArticleListResult>("/article/page", {
                params: {
                    "page-current": pageCurrent,
                    "page-size": pageSize,
                    "tag-ids": tagIds,
                }
            }).then(res => res.data);
        }
    }

    export class Category {
        public static Add = (dto: CategoryAddDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/category/add", dto).then(res => res.data);
        }

        public static Delete = (dto: CategoryDeleteDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/category/delete", dto).then(res => res.data);
        }

        public static GetList = (): Promise<CategoryListResult> => {
            return axiosInstance.get<CategoryListResult>("/category/list").then(res => res.data);
        }
    }

    export class Tag {
        public static Add = (dto: TagAddDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/tag/add", dto).then(res => res.data);
        }

        public static Delete = (dto: TagDeleteDTO): Promise<EmptyResult> => {
            return axiosInstance.post<EmptyResult>("/tag/delete", dto).then(res => res.data);
        }

        public static GetList = (): Promise<TagListResult> => {
            return axiosInstance.get<TagListResult>("/tag/list").then(res => res.data);
        }
    }
}

export default API;
