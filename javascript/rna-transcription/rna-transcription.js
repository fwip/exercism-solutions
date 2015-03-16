module.exports = function(){
  var comp = {A: "U", C: "G", T: "A", G: "C"};
  return function(dna){
    return dna.split("").map(function(n){ return comp[n]; }).join("");
  };
}();
