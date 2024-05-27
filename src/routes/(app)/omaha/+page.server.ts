export async function load({locals}) {
    return {
        auth: locals.userPb.authStore.token,
    }
}


