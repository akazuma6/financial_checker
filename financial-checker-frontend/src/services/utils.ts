const baseUrl = process.env.NEXT_PUBLIC_API_BASE_URL || 'http://localhost:8080';

export interface FetchOptions {
  url: string;
  method: 'GET' | 'POST' | 'PUT' | 'DELETE';
  body?: unknown;
  query?: Record<string, string | number | undefined>;
}

export async function fetchWithQuery<T>({
  url,
  method,
  query,
}: FetchOptions): Promise<T> {
  const queryString = query
    ? '?' +
      Object.entries(query)
        .filter(([_, value]) => value !== undefined)
        .map(([key, value]) => `${key}=${encodeURIComponent(String(value))}`)
        .join('&')
    : '';

  const response = await fetch(`${url}${queryString}`, {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
}

export async function fetchWithBody<T>({
  url,
  method,
  body,
}: FetchOptions): Promise<T> {
  const response = await fetch(url, {
    method,
    headers: {
      'Content-Type': 'application/json',
    },
    body: body ? JSON.stringify(body) : undefined,
  });

  if (!response.ok) {
    throw new Error(`HTTP error! status: ${response.status}`);
  }

  return response.json();
}

export { baseUrl };
