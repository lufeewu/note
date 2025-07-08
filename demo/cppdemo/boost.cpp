#include <iostream>
#include <string>
#include <boost/regex.hpp>

int main() {
  std::string text = "This is a test string.";
  boost::regex pattern("test");

  if (boost::regex_search(text, pattern)) {
    std::cout << "Pattern found!" << std::endl;
  }

  return 0;
}