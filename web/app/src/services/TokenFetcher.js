export async function fetchToken(username, token) {
    const data = {
        username: username
    }
    const response = await fetch("http://localhost:8080/api/v1/auth/token", {
        method: "POST",
        headers: new Headers(
            {
                'content-type': 'application/json',
                'Authorization': `Bearer ${token}`
            }
        ),
        body: JSON.stringify(data)
    });
    let jResp = await response.json();
    // TODO: not store token in local storage
    let tokenWrapper = jResp.token
    localStorage.setItem('token', tokenWrapper['token'])
    return jResp;
}

export function decodeToken(token) {
    let base64Url = token.split('.')[1];
    let base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    let jsonPayload = decodeURIComponent(window.atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
}