// routes/signup/+page.server.ts
//import { lucia } from "../../auth";
import { redirect } from "@sveltejs/kit";
import { hash } from "@node-rs/argon2";
import { prisma } from "$lib/prisma"
//import { v4 as uuidv4 } from "uuid";

import type { Actions } from "./$types";

export const actions: Actions = {
	default: async ({request}) => {
        
        const data = await request.formData();

        let username = data.get("username") as string
        let password = data.get("password") as string

		//const userId = uuidv4(); // 16 characters long
		const passwordHash = await hash(password, {
			// recommended minimum parameters
			memoryCost: 19456,
			timeCost: 2,
			outputLen: 32,
			parallelism: 1
		});

		// TODO: check if username is already used
		await prisma.users.create({
            data: {
			    username: username,
			    password: passwordHash,
            }
		});

		redirect(302, "/");
	}
};
