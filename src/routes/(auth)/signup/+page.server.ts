import { redirect } from "@sveltejs/kit";
import PocketBase from 'pocketbase';
const pb = new PocketBase('http://127.0.0.1:8090');
export const actions = {
    default: async ({ request }) => {

        const data = await request.formData();

        let username = data.get("username") as string
        let password = data.get("password") as string
        let email = data.get("email") as string;
        // TODO: check if username is already used
         
        const payload = {
            "username": username,
            "password": password,
            "passwordConfirm": password,
            "email": email,
        };
        console.log(payload);
        
        const authData = await pb.collection('users').create( payload );
        console.log(authData); 
        redirect(302, "/");
    }
};
