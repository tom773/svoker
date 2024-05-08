import { redirect } from "@sveltejs/kit";
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
export const actions = {
    default: async ({ request }) => {

        const data = await request.formData();

        let password = data.get("password") as string
        let email = data.get("email") as string;
        // TODO: check if username is already used
         
        const authData = await pb.collection('users').authWithPassword(email, password);
        console.log(authData);
        console.log(pb.authStore.isValid);
        console.log(pb.authStore.token);
        console.log(pb.authStore.model.id);
        redirect(302, "/");
    }
};
