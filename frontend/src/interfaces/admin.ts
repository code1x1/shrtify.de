export interface PostRequest<T> {
    request: T;
    meta: unknown;
}

export interface PostResponse<T> {
    response: T;
    meta: unknown;
}
