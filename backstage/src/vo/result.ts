export interface Result<T> {
    hasError: boolean;
    error: string;
    message: string;
    data: T | null;
}