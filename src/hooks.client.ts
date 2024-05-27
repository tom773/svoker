import { pb } from '$lib/pocketbase';

pb.authStore.loadFromCookie(document.cookie);
pb.authStore.onChange(() => {
    pb.authStore.exportToCookie({httpOnly: false});
}, true); 
