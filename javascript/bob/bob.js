//
// This is only a SKELETON file for the "Bob" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Bob = function() {};

Bob.prototype.hey = function(input) {
  // Bob gets upset if you say only whitespace
  if (input.trim() == ""){
    return "Fine. Be that way!";
  }

  // Bob's mellow gets harshed if you're saying only uppercase letters
  if (input == input.toUpperCase() && input != input.toLowerCase()){
    return "Whoa, chill out!";
  }

  // Bob goes with the flow if a question-mark ends the query
  if (input.indexOf("?") == input.length - 1){
    return "Sure.";
  }

  // Bob is noncommittal
  return "Whatever.";
};

module.exports = Bob;
