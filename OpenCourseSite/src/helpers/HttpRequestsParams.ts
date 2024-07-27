const _params = (body?: string, headers?: HeadersInit): RequestInit => ({
  credentials: "include",
  mode: "cors",
  headers: headers ? headers : { "Content-Type": "application/json" },
  body,
});

export const getParams: RequestInit = {
  ..._params(),
  method: "GET",
};

export const postParams = (
  body?: string,
  headers?: HeadersInit
): RequestInit => ({
  ..._params(body, headers),
  method: "POST",
});

export const putParams = (
  body?: string,
  headers?: HeadersInit
): RequestInit => ({
  ..._params(body, headers),
  method: "PUT",
});

export const deleteParams = (
  body?: string,
  headers?: HeadersInit
): RequestInit => ({
  ..._params(body, headers),
  method: "DELETE",
});
