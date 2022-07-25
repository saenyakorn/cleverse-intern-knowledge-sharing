self.onmessage = function (e) {
  function fibonacci(n) {
    if (n <= 1) {
      return n
    }
    return fibonacci(n - 1) + fibonacci(n - 2)
  }

  const counter = e.data
  const result = fibonacci(40) + counter
  console.log('RESULT = ' + result)
  self.postMessage(result)
}
