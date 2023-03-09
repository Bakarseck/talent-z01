#ifndef HEADER_H
#define HEADER_H

#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Date
{
    int month;
    int day;
    int year;
};

// all fields for each record of an account
struct Record
{
    int id;
    int userId;
    char name[100];
    char country[100];
    int phone;
    char accountType[10];
    int accountNbr;
    double amount;
    struct Date deposit;
    struct Date withdraw;
};

// manage the amount
double deposit(double amountDeposited, Record record) ;
double withDraw(double amountWithDrawed, Record record) ;

struct User
{
    int id;
    char name[50];
    char password[50];
};

// typeDef the structs
typedef struct User User ;
typedef struct Date Date ;
typedef struct Record Record ;

// authentication functions
void loginMenu(char a[50], char pass[50]);
void registerMenu(char a[50], char pass[50]);
const char *getPassword(User newUser);

// The functions of the managing atm-managing
void registerNewUser() ;
void updateInformation() ;
void updateCountry(int id) ;
void checkingDetailsAccount(int id) ;


// system function
void createNewAcc(User newUser);
void mainMenu(User newUser);
void checkAllAccounts(User newUser);

#endif
