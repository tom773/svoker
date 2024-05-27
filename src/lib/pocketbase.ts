import PocketBase from 'pocketbase';

export function newInstance(){
    return new PocketBase("http://localhost:8090");
}
export const pb = newInstance();
