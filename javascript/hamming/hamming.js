module.exports.compute = function(a, b){

  if (a.length != b.length) {
    throw Error("DNA strands must be of equal length.");
  }

  // Naive approach, step through and compare each character
  var diff = 0;
  for (var i=0; i<a.length; i++){
    if (a[i] != b[i]){
      diff ++;
    }
  }
  return diff;
};
