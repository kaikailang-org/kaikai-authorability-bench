pub struct Account {
    pub id: i64,
    pub balance: i64,
}

fn find(ledger: &[Account], id: i64) -> Option<usize> {
    ledger.iter().position(|a| a.id == id)
}

fn clone_l(ledger: &[Account]) -> Vec<Account> {
    ledger.iter().map(|a| Account { id: a.id, balance: a.balance }).collect()
}

pub fn open_account(ledger: &[Account], id: i64) -> Option<Vec<Account>> {
    if find(ledger, id).is_some() {
        return None;
    }
    let mut out = clone_l(ledger);
    out.push(Account { id, balance: 0 });
    Some(out)
}

pub fn deposit(ledger: &[Account], id: i64, n: i64) -> Option<Vec<Account>> {
    let i = find(ledger, id)?;
    if n < 1 {
        return None;
    }
    let mut out = clone_l(ledger);
    out[i].balance += n;
    Some(out)
}

pub fn withdraw(ledger: &[Account], id: i64, n: i64) -> Option<Vec<Account>> {
    let i = find(ledger, id)?;
    if n < 1 || ledger[i].balance < n {
        return None;
    }
    let mut out = clone_l(ledger);
    out[i].balance -= n;
    Some(out)
}

pub fn transfer(ledger: &[Account], from: i64, to: i64, n: i64) -> Option<Vec<Account>> {
    let fi = find(ledger, from)?;
    let ti = find(ledger, to)?;
    if n < 1 || from == to || ledger[fi].balance < n {
        return None;
    }
    let mut out = clone_l(ledger);
    out[fi].balance -= n;
    out[ti].balance += n;
    Some(out)
}

pub fn balance_of(ledger: &[Account], id: i64) -> Option<i64> {
    Some(ledger[find(ledger, id)?].balance)
}
