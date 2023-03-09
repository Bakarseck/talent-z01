#include "header.h"

// manage the amount
double deposit(double amountDeposited, User newUser) ;
double withDraw(double amountWithDrawed, User newUser) ;

// authentication functions
void loginMenu(char a[50], char pass[50]);
void registerMenu(char a[50], char pass[50]);
const char *getPassword(User newUser);

// The functions of the managing atm-managing
void registerNewUser() ;
void updatePhoneNumber(int id) ;
void updateCountry(int id) ;
void checkingDetailsAccount(int id) ;
void makeTransaction(double amountTransacted, User newUser) ;
void 


// system function
void createNewAcc(User newUser);
void mainMenu(User newUser);
void checkAllAccounts(User newUser);