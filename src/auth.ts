import { Lucia } from "lucia";
import { dev } from "$app/environment";
import { PrismaAdapter } from "@lucia-auth/adapter-prisma";
import { PrismaClient } from "@prisma/client";
import { prisma } from "$lib/prisma";

const client = new PrismaClient();
const adapter = new PrismaAdapter(client.session, client.users);

const users = await prisma.users.findMany();
console.log(users);

export const lucia = new Lucia(adapter, {
	sessionCookie: {
		attributes: {
			secure: !dev,
		}
	},
    getUserAttributes: (attributes) =>{
        return{
            username: attributes.username,
            password: attributes.password,
        };
    },
});

interface DatabaseUserAttributes {
	username: string;
    password: string;
}

declare module "lucia" {
	interface Register {
		Lucia: typeof lucia;
        DatabaseUserAttributes: DatabaseUserAttributes;
	}
}
