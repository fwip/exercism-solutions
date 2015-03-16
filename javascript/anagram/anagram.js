function anagram(match){
  this.match = match;
}

function sortString(s){
  return s.toLowerCase().split('').sort().join('');
}

anagram.prototype.matches = function(candidates){
  // This is ugly - there's surely a cleaner way to handle string arguments
  if (typeof(candidates) === "string"){
    candidates = Array.prototype.slice.call(arguments);
  }

  var match = this.match;
  return candidates.filter(function(e){
    return e.toLowerCase() !== match.toLowerCase() &&
           sortString(e) === sortString(match);
  });
};

module.exports = function(match){
  return new anagram(match);
};
