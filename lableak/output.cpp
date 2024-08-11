#include <iostream>
using namespace std;

class BankAccount {
private:
    string name;
    string acc_no;
    double balance;

public:
	BankAccount() {}
    BankAccount(string n, string acc, double bal) : name(n), acc_no(acc), balance(bal) {}

    void printInfo() const;
    void read();
    void deposit();
    void withdraw();
};

void BankAccount::printInfo() const {
    cout << "Account Holder: " << name << endl;
    cout << "Account Number: " << acc_no << endl;
    cout << "Balance: Rs." << balance << endl;
}

void BankAccount::read() {
    cout << "Enter account holder name: ";
    getline(cin, name);
    cout << "Enter account number: ";
    getline(cin, acc_no);
    cout << "Enter initial balance: Rs.";
    cin >> balance;
}

void BankAccount::deposit() {
    double amount;
    cout << "Enter deposit amount: Rs.";
    cin >> amount;

    if (amount > 0) {
        balance += amount;
        cout << "Deposit successful. New balance: Rs." << balance << endl;
    } else {
        cout << "Invalid deposit amount." << endl;
    }
}

void BankAccount::withdraw() {
    double amount;
    cout << "Enter withdrawal amount: Rs.";
    cin >> amount;

    if (amount > 0 && amount <= balance && balance - amount >= 500) {
        balance -= amount;
        cout << "Withdrawal successful. New balance: Rs." << balance << endl;
    } else if (amount > balance || balance - amount < 500) {
        cout << "Insufficient funds." << endl;
    } else {
        cout << "Invalid withdrawal amount." << endl;
    }
}

int main() {
    BankAccount account;
	account.read();
	//Menu driven operations.
	cout << "Bank Operations" << endl;
	cout << "---------------" << endl;
	cout << "1. Display." << endl;
	cout << "2. Deposit." << endl;
	cout << "3. Withraw." << endl;
	cout << "4. Exit." << endl;
	int choice;
	while(true) {
		cin >> choice;
		if(choice == 1) {
			account.printInfo();
		} else if(choice == 2) {
			account.deposit();
		} else if(choice == 3) {
			account.withdraw();
		} else if(choice == 4) {
			cout << "Thank you for banking with Us." << endl;
			break;
		} else{
			cout << "Invalid choice. Try Again." << endl;
		}
	}

    return 0;
}
