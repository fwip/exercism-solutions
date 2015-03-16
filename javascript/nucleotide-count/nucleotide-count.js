module.exports = function(dnaInit){
  // If given no input, use the empty string
  var dna = (dnaInit === undefined) ? "" : dnaInit;

  // Test for invalid input
  if (/[^ACTG]/.test(dna)){
    throw new Error("DNA is only A, C, T, and G. Doofus.");
  }

  // Calculate the histogram now
  var histo = {A:0, G:0, T:0, C:0};
  dna.split('').forEach(function(c) {
    histo[c]++;
  });

  return {
    count: function(c){ return histo[c]; },
    histogram: function(){ return histo; },
  };
};
