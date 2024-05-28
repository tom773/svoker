import { redirect, error } from '@sveltejs/kit';
export const load = ({ locals }) => {
	if (!locals.userPb.authStore.isValid) {
		throw redirect(303, '/signin');
	} else {
        console.log("authenticated"); 
    } 
};

export const actions = {
    username: async ({ request, locals }) => {
        const body = await request.formData();
        try {
            const { username } = await locals.userPb.collection('users').update(locals?.user?.id, body);
            locals.user.username = username;
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, '/settings')
    },
    password: async ({ request, locals }) => {
        const body = await request.formData();
        try {
            await locals.userPb.collection('users').update(locals?.user?.id, {"password": body.get('password')});
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, '/settings')
    },
    email: async ({ request, locals }) => {
        const body = await request.formData();
        try {
            const { email } = await locals.userPb.collection('users').requestEmailChange(body.get('email'));
            locals.user.email = email;
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, '/settings')
    },
    amt: async ({ request, locals }) => {
        const body = await request.formData();
        try {
            let bal = parseInt(locals.user.balance);
            let amt = parseInt(body.get("amt"));
            const { balance } = await locals.userPb.collection('users').update(locals?.user?.id, {"balance": (bal+amt)});
            locals.user.balance = balance;
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }
        throw redirect(303, '/settings')
    },
    avatar: async ({ request, locals }) => {
        const body = await request.formData();
        try {
            const { avatar } = await locals.userPb.collection('users').update(locals?.user?.id, {"avatar": body.get('avatar')});
            locals.user.avatar = avatar;    
        } catch (err) {
            console.log(err);
            throw error(500, 'You probably made a typo');
        }

        return {
            success: true
        }
    }
    
};

