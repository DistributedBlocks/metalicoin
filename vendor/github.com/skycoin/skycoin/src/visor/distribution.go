package visor

import (
	"github.com/skycoin/skycoin/src/coin"
)

const (
	// MaxCoinSupply is the maximum supply of skycoins
	MaxCoinSupply uint64 = 3e8 // 300,000,000 million

	// DistributionAddressesTotal is the number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	// DistributionAddressInitialBalance is the initial balance of each distribution address
	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// InitialUnlockedCount is the initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100

	// UnlockAddressRate is the number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 0

	// UnlockTimeInterval is the distribution address unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// GetDistributionAddresses returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetUnlockedDistributionAddresses returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (25) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (25) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (25) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// GetLockedDistributionAddresses returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// TransactionIsLocked returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"2PF2Y8rHN7Db97B4vro7Ug6yKx9SzsnF5Kb",
"2f5LFnckrKnT5x7EeyLzSQm5kdm5xmvXo16",
"JGFMJWbL5FVv3xuan4W3Cv3DeUa4cprtUm",
"9qbTGBG8LpsUeGgw195RkEeAJi8W4oJ79h",
"9isU4hwri7yfFirMAVHJ4hB3WPbnmvGw4Z",
"9o9BUTZ5jEtunXEfMDddSqZp6XB1UkTyrz",
"Yf7a84bxunuFVExHybrohPwXRVVMKsy5Dd",
"sw8tG5ZkSvvLg3UcmZMcGNomhhkZ1gVqQP",
"usgwWr1ewu3PuW2rXpx7HFwSiAuaKpXyK5",
"Q5rTZMxGoZJeMyuhBCkQEiR3G6PssgLLDg",
"2JnJR9iXVyTSLuB6uSu9Hjr8b8sasA54JTx",
"2Q9LJ3GQNZCUfEQLw5CUiTZVcN7hTjgjqWe",
"bYS77oNLNQrnKV6cxq3mzdVN8xCWpHfp2N",
"2cYGDJYYmzpjoGwDnSCDonCcHXjMTsRXeGs",
"gYfzNsWCw7kBXrrrRYrMjGuWSVfVNhRz7g",
"K1iCkJ1PB91PCny8xmGsXQcX6Qseb8fJdm",
"UebaLXSKLYk8bEPkCo7ftuZVyKDhLrC9PP",
"2Wnj12PTdsECiRnQrjBngwUsU4PySJEDT7R",
"2HK2Ap7ponx6Q7BY2VPPXoDnu944Uh4tExs",
"perNzwS1sR3hzQtmS7Q3a9oLa3LnTt6US1",
"2CR7AUfPuWDD41Lt3qRGx5UGPVkLVrSVnSt",
"2AQKURcAAuotHwQ28VoRkB3LkMDeVWsWjCZ",
"S95PCDcH44QNVGvzvHxBKqq35zgGCyaVHN",
"VAb5KgWJCTNhSsUHPwqvMBdfXs8tehEmKP",
"z572fLQ3rjFdf7SaMRcJvSi3zsSRx2Qapv",
"2JbySKAbQ8DhoYtpBTY5YJbg2LX9eQppLYm",
"2XoBYF9wEFB6T5VB834QtAAZrZBdpk6RA8o",
"21GvuwQkGv4WLgZCoJpwQMnQCkGNWQfd11y",
"vvJZhZD3mELnvpR9wznbrBZpSozt5MdV12",
"Z7dYmXZPuPxxxFvuJ5dfbdo1eVFXYSXyvY",
"h3rm1V797Za6cqguGeSpv8hQtqfKLa3qP5",
"4UkziCJMiz5yWMkg8rajkeM2wH9myakqfM",
"KSb4yKyKNn9j5Z51jNE2M1yXvgBPT825oa",
"2jtgrzBui3otA1UFyx1s5UTmKQYVaaiK5gq",
"28WEM7DKbV7ijSPvRAwGmPGGJix8Ri3z2bt",
"2QVEZ3AS6KrRx9WkiWFRK7rKMLGfQuCWYz9",
"2GQxPXR37hXLA5EBRcjRYkKew9ZPZWSaapc",
"2Ew7NvNAFNffVXUTyy8ysyEBaFkmENqKA2s",
"2iB91tgsFQKPRxKZLNf4DHXenwtZeJs6ze",
"2DqkrngTepqjVmpZhufNs6mC3KdPgGd3gDq",
"Z4tGRjiR1csxXb6Nm7Xx1tHJnWun1haHTj",
"He88dCdfYzcqRX7a9i4kBj4VAFMLko1X2c",
"zGHpytg6hJe4T81oxu9CeDTLHQZc5jRsCt",
"23gKeuu8zUWYH3gE7Ut4LWNqnRRWSHTWUQv",
"271CMFaMBdUXJnzTJqrRtsPyBC6P4SfVqdB",
"24GYbKSD4StrPF2jJXRtNsXRkGQnQFSVQvz",
"29JKT26nCGPSzdiYFRMKNpkTvFfsiq9SZPV",
"Y1moyqwjwSeGw4B1X5SLdsP2omMqhZ3e2U",
"5UoSbq7XSt3eg66dWb19K4A9xZXtFCCs1F",
"DpWmiQdQauWtngVus7JGstG3N3F5sqGPq1",
"2hJg3AQAV3fzQkySCsBaHNKg1Mt7j1HHMgG",
"rhnGEed5qWhYXHtfve3LwGYzhtZD7V5yWd",
"95vbjuRLk527GNgeuqZ4UX6KhS955aULbS",
"21sMY48dihceeAZDdWQfaWyesksyMHwXadZ",
"Ffx9R67Jm7a6UUuGWmrNWdFpTC7mhe62T1",
"27kpPnkNKyzGKu2X5urNyh2mCXxUmti8G2b",
"gqriLTshb9tauj7WgCmE5rGoxYkkyEwRer",
"ns6ymZb1C67GGoADTrjWRWaMswwhhYJBwC",
"2Kus54feSmVWunnRNKNNYzH6MxcXt9NjNoX",
"7KkrPTk3tmbNPKaH6FwP6upY3HktXVEATE",
"2c2baiCgT3i4SYw3ox3GuCaWs14pLmw8xwK",
"2bjiFRQskc1spGRWP21jY6psPVtWQMnGDzd",
"GyoZmRatGc82VZLfATgNMjXCoJmjQZTjdL",
"VAtDa9NnBnq7xr6hf3DcJjdvYbqYeThEGX",
"2GHe6W4E4UDHE1fqBe8KDmL5qETLzVyhAhR",
"2NnHehm9ZFkGts3zPKUf9HJ8ceT74DsMxnd",
"2VneHrjCgaGoBizHbQcQUCvK5u9TRL5u88Z",
"SHL6Ca6pWTxqJ3bpJSXDGVEpgYY3fcqoYi",
"2VB6DfEtLsqe8jkkFsDCSY7u7FCsnKc8dkr",
"KehUdaPQ19z2Vhp4GSfZfrWChKuibGHbhV",
"2d6z8BEQerSDWbyFNPaq4KbEz5abG9JxE8D",
"pfPphAixLJ2qXSZisvCNgUfcRxJaPwJf5i",
"rrLGYH2mor2arh3nXecXmUyqfzkXgf7BoE",
"YRxJPpqC7EFChYv7sFJpvoe58qjtsPWsCT",
"CSmzt4qp1YDCqMBp8oKa78rYatoTJQbi7d",
"YyC4amgwmSXCrQN6n4bGFky8yyDtXBPxg6",
"tShXfL659oe1ZwcxVo4PZyeq5rPvUkYBSz",
"XeBknaBiC3HCnqkUQC52jjv5QmGiGn6NDz",
"2HhFub68y49cS7FhxW93p7gHjPnLFAwty3A",
"2U3RhjNiDeB4wcVcjgXrts7i9i7VpWfbVjW",
"2WUL9AS16b1A6D7y4w1JC2ndUPeEX3EwFy9",
"uAKK2Q5WTzfGAL7mdFfFiueRCijxaoQNrN",
"gT7mXuzAQQTdfDzsCjWpPGcVRCSZqnumCB",
"LthmQSTNL1Yuw8bXQ6BAuMa74JQnU6MTtB",
"XPL8viR4J2mYi5MEzC2JEaUbQDV5KSvwex",
"29rkMoWA1RMrhJd6byy8b5tWXjwJJ4PA3vB",
"GsBqFcuecCZyKAmqmZDUZnc5nZFG3nRjNx",
"SL822CmXu5pMqdDuzUHAMNwL292HqDMEkT",
"PJbiLKGz4ao4TdcU9r43e1f9kayjt81oRG",
"BCPb2QFaEXArmAcaHZ7YVpFd1Chv6sXs86",
"X5D6DZ9acPJURbDfN3tTDbc4t7UBXhydy6",
"Eg3nZsPLpxjm3wwgrsWvhC5aAfu2EwqQKX",
"2GcNfw6TwiE5pVDN8YFxpmxdGqcN86H2mvp",
"2hrZx9SnXQEfouce5yXSM3YfXAPsGMnVa26",
"Ps34Ey1yTGxTZfSaWY8XiiStQJxchVnuuz",
"2bFHyt7XTUh6wDdHE1bfoEZ8MXKK2ppU6h",
"cYN9jB3Y24uS4TV31KucRmZPwaohfxPSQ9",
"2F9SDGWjbARK8ZU9qyKNPxK931SMcCiXjQq",
"2Nw3qP6VjuWuUZwcfkaiXuSPzZ4CT73gZX5",
"gvHoKASwyeXZKLSj5wjZHLGpEJmJ7ascpZ",
}
