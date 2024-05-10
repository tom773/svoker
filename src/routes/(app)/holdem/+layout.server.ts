import PocketBase from 'pocketbase';

const pb = new PocketBase('http://127.0.0.1:8090');

const resultList = await pb.collection('tables').getFullList({
    sort: 'tnum',
});

export const load = ({ locals }) => {
    locals.pb = pb;
    locals.resultList = resultList;
}


