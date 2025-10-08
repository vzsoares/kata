function solution(arr){
  const mem = {};
  const size  = arr.length
  
  for (let i = 0; i < size -1; i++){
    for (let j = i; i < size - 1;){
      mem[arr[i]] = 1
      if (arr[i] === arr[j] + 1){ 
        mem[arr[i]] = mem[arr[i]] + 1
      }
        j++
    }
  }
  
  return Object.keys(mem).sort(1)[0]
}

function solution2(arr){ // O(nlogn)
  const sorted = arr.sort(); // {1, 2, 3, 4, 9, 10, 20} -> 1, 3, 4, 2 -> 4
  let sub = 0;
  let prev = arr[0];
  const res = [];
  
  for (let i = 1; i < arr.length -1; i++) {
    if (prev === arr[i] + 1){
      res[sub] = (res[sub]??0) + 1
    } else{
        sub = sub + 1
    }
    prev = arr[i]
  }
  
  return res.sort()[res.length-1]
  
}

function solutionBest(arr){
  const set = new Set(arr);
  let max = 0;
  
  for (const num of set){ // O(n)
    if (!set.has(num-1)) {
      let current = num;
      let len = 1;
      while (set.has(cur + 1)){
        cur++;
        len++;
        if (len > max) max = len;
      }
    }
  }
  return max;
}
