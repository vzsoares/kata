/*
Implement an async function `retry(fn, retries = 3, delay = 200)` that calls an async `fn()`.
- If `fn` rejects, retry up to `retries` times with `delay` ms between attempts.
- On final failure, rethrow the last error.
Example:
  await retry(() => fetchJSON(url));
*/

function sleep(ms: number) {
    return new Promise((res) => setTimeout(res, ms));
}

async function retry(fn: () => Promise<unknown>, retries = 3, delay = 200) {
    for (let i = 0; i < retries; i++) {
        try {
            console.log(i);
            const res = await fn();
            console.log('success');
            return res;
        } catch (e) {
            console.log('error');
            if (i === retries) throw e;
            await sleep(delay);
        }
    }
}

console.log(
    retry(() => new Promise((r, e) => (Math.random() > 0.4 ? r(true) : e()))),
);
