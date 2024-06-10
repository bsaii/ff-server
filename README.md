# FinanceForge

FinanceForge is a server in Go with documentation in Swagger that lets you interact with a smart contract deployed on the Sepolia testnet. It provides data from events emitted by the smart contract. FinanceForge is a basic staking contract with lending and borrowing functionality.

## Features

### Functions

- **stake:** Allows users to stake tokens. Updates the user's staking balance and emits the `Staked` event.
- **withdraw:** Allows users to withdraw staked tokens. Updates the user's staking balance and emits the `Withdrawn` event.
- **claimReward:** Allows users to claim their accumulated staking rewards. Updates the user's rewards balance and emits the `RewardPaid` event.
- **earned:** Calculates the rewards earned by a user based on the staked time and reward rate.
- **lend:** Allows users to lend tokens to the contract. Updates the user's lending balance and emits the `Lent` event.
- **borrow:** Allows users to borrow tokens by providing collateral. Updates the user's borrowing and collateral balances, and emits the `Borrowed` event.
- **repay:** Allows users to repay borrowed tokens. Updates the user's borrowing and collateral balances, and emits the `Repaid` event.

### State Variables

- **stakingToken:** The ERC20 token used for staking, lending, and borrowing.
- **rewardRate:** The rate at which staking rewards are calculated.
- **interestRate:** The rate at which interest on loans is calculated.
- **users:** Maps user addresses to their staking, lending, borrowing, rewards, and collateral balances.

### Events

- **Staked:** Emitted when a user stakes tokens.
- **Withdrawn:** Emitted when a user withdraws staked tokens.
- **RewardPaid:** Emitted when a user claims their rewards.
- **Lent:** Emitted when a user lends tokens.
- **Borrowed:** Emitted when a user borrows tokens.
- **Repaid:** Emitted when a user repays borrowed tokens.

## Smart Contract Interaction

The contract has been deployed on the Sepolia testnet. You can interact with the smart contract using the following links:

- **Swagger Documentation:** [Swagger Link](http://localhost:8000/swagger/index.html)
- **Read from Contract:**
  - [Read FF Contract Link](https://sepolia.etherscan.io/address/0x1cD752eD8131e62f221D4EdF1Ac48d101933f495#readContract)
  - [Read FF Token Contract Link](https://sepolia.etherscan.io/address/0xFfB5e8Cc951554f51140960d87d16598945B226E#readContract)
- **Write to Contract:**
  - [Write FF Contract Link](https://sepolia.etherscan.io/address/0x1cD752eD8131e62f221D4EdF1Ac48d101933f495#writeContract)
  - [Write FF Token Contract Link](https://sepolia.etherscan.io/address/0xFfB5e8Cc951554f51140960d87d16598945B226E#writeContract)

## Getting Started

To get started with the FinanceForge server, follow these steps:

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. **Clone the repository:**

    ```sh
    git clone https://github.com/yourusername/financeforge.git
    cd financeforge
    ```

2. **Install dependencies:**

    ```sh
    go mod tidy
    ```

3. **Create and configure the `.env` file:**

    ```sh
    touch .env
    ```

    Add the following environment variables to the `.env` file:

    ```
    PORT=port
    ETH_NODE_URL=websocket_url
    DB_PATH=database_file_path
    FF_ADDRESS="0x1cD752eD8131e62f221D4EdF1Ac48d101933f495"
    FFT_ADDRESS="0xFfB5e8Cc951554f51140960d87d16598945B226E"
    ```

4. **Build the application:**

    ```sh
    go build -o financeforge
    ```

5. **Run the application:**

    ```sh
    ./financeforge
    ```

Now, you can access the server and interact with the smart contract using the provided endpoints.

## Build File

You can also download the pre-built binary from the repository's releases and run the app directly:

- [Download Build File](link-to-build-file) - not available

