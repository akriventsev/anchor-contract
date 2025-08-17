# First Contract - Solana Program Documentation

## Overview

This is a Solana program built using the Anchor framework. The program provides two main instruction functions for demonstration purposes and includes a JavaScript client for interaction.

## Project Structure

```
├── programs/
│   └── first-contract/          # Solana program source code
│       ├── src/
│       │   └── lib.rs          # Main program logic
│       └── Cargo.toml          # Rust dependencies
├── tests/                       # Test files
│   └── first-contract.js       # JavaScript tests
├── client.js                   # JavaScript client example
├── Anchor.toml                 # Anchor configuration
├── Cargo.toml                  # Workspace configuration
└── package.json                # Node.js dependencies
```

## Solana Program API

### Program ID
```
Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS
```

### Public Functions

#### 1. `initialize()`

**Description**: A simple instruction that logs "Karamba" to the program log.

**Context**: `Initialize` (no account validation required)

**Returns**: `ProgramResult`

**Example Usage**:
```rust
// Rust
pub fn initialize(ctx: Context<Initialize>) -> ProgramResult {
    msg!("Karamba");
    Ok(())
}
```

```javascript
// JavaScript
const tx = await program.rpc.initialize();
console.log("Transaction signature:", tx);
```

#### 2. `initialize2()`

**Description**: A more complex instruction that demonstrates program address derivation and account manipulation.

**Context**: `Initialize2` (no account validation required)

**Returns**: `ProgramResult`

**Functionality**:
- Decodes a base58-encoded public key
- Creates a program-derived address (PDA)
- Demonstrates account info manipulation
- Logs "Karamba2" to the program log

**Example Usage**:
```rust
// Rust
pub fn initialize2(ctx: Context<Initialize2>) -> ProgramResult {
    // Decode base58 address
    let decoded = bs58::decode("2eGVwJcrrEWResAeRBB6AXrksgN1WHss9eWUUQgXz8Ux").into_vec().unwrap();
    let addr = &Pubkey::new(decoded.as_slice());
    
    // Create program-derived address
    let seed = bs58::decode("2eGVwJc").into_vec().unwrap();
    let s: Vec<&[u8]> = vec![seed.as_slice()];
    let seed_slice = s.as_slice();
    let new_addr = Pubkey::create_program_address(seed_slice, &addr)?;
    
    msg!("Karamba2");
    Ok(())
}
```

```javascript
// JavaScript
const tx = await program.rpc.initialize2();
console.log("Transaction signature:", tx);
```

### Account Structures

#### `Initialize`
```rust
#[derive(Accounts)]
pub struct Initialize {
    // No account validation required
}
```

#### `Initialize2`
```rust
#[derive(Accounts)]
pub struct Initialize2 {
    // No account validation required
}
```

## JavaScript Client API

### Setup and Configuration

#### 1. Provider Configuration
```javascript
const anchor = require('@project-serum/anchor');

// Configure for local cluster
anchor.setProvider(anchor.Provider.local('https://api.devnet.solana.com', anchor.Provider.defaultOptions()));

// Configure for environment (recommended for testing)
anchor.setProvider(anchor.Provider.env());
```

#### 2. Program Client Generation
```javascript
// Read the generated IDL
const idl = JSON.parse(require('fs').readFileSync('./target/idl/first_contract.json', 'utf8'));

// Program ID (update this with your deployed program address)
const programId = new anchor.web3.PublicKey('YOUR_PROGRAM_ID_HERE');

// Generate the program client
const program = new anchor.Program(idl, programId);
```

### Client Methods

#### `program.rpc.initialize()`
Executes the `initialize` instruction.

**Returns**: `Promise<string>` - Transaction signature

**Example**:
```javascript
try {
    const tx = await program.rpc.initialize();
    console.log("Initialize transaction:", tx);
} catch (error) {
    console.error("Error:", error);
}
```

#### `program.rpc.initialize2()`
Executes the `initialize2` instruction.

**Returns**: `Promise<string>` - Transaction signature

**Example**:
```javascript
try {
    const tx = await program.rpc.initialize2();
    console.log("Initialize2 transaction:", tx);
} catch (error) {
    console.error("Error:", error);
}
```

## Testing

### Running Tests
```bash
# Run all tests
anchor test

# Run specific test file
anchor test tests/first-contract.js
```

### Test Examples

#### Basic Test Structure
```javascript
const anchor = require('@project-serum/anchor');

describe('first-contract', () => {
    // Configure the client to use the local cluster
    anchor.setProvider(anchor.Provider.env());

    it('Should initialize successfully', async () => {
        const program = anchor.workspace.FirstContract;
        const tx = await program.rpc.initialize();
        console.log("Transaction signature:", tx);
    });

    it('Should initialize2 successfully', async () => {
        const program = anchor.workspace.FirstContract;
        const tx = await program.rpc.initialize2();
        console.log("Transaction signature:", tx);
    });
});
```

## Building and Deployment

### Build the Program
```bash
# Build the program
anchor build

# Build and generate IDL
anchor build --skip-lint
```

### Deploy to Local Cluster
```bash
# Start local validator
solana-test-validator

# Deploy program
anchor deploy
```

### Deploy to Devnet
```bash
# Switch to devnet
solana config set --url devnet

# Deploy program
anchor deploy --provider.cluster devnet
```

## Dependencies

### Rust Dependencies
- `anchor-lang = "0.17.0"` - Anchor framework
- `bs58 = "0.4.0"` - Base58 encoding/decoding

### JavaScript Dependencies
- `@project-serum/anchor = "^0.17.0"` - Anchor JavaScript client
- `chai = "^4.3.4"` - Testing assertion library
- `mocha = "^9.0.3"` - Testing framework

## Configuration Files

### Anchor.toml
```toml
[features]
seeds = false
skip-lint = false

[programs.localnet]
first_contract = "Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS"

[registry]
url = "https://api.apr.dev"

[provider]
cluster = "localnet"
wallet = "~/.config/solana/id.json"
```

### Cargo.toml (Workspace)
```toml
[workspace]
members = [
    "programs/*"
]
```

## Error Handling

### Common Error Types
- `ProgramResult::Err` - Program execution errors
- `AnchorError` - Anchor-specific errors
- `SolanaError` - Solana blockchain errors

### Error Handling Example
```javascript
try {
    const tx = await program.rpc.initialize();
    console.log("Success:", tx);
} catch (error) {
    if (error.error?.errorCode) {
        console.error("Anchor Error:", error.error.errorCode);
    } else {
        console.error("General Error:", error);
    }
}
```

## Security Considerations

1. **Program ID**: Always verify the program ID matches your deployed program
2. **Account Validation**: The current implementation doesn't validate accounts - consider adding proper validation for production use
3. **Error Handling**: Implement proper error handling in production applications
4. **Testing**: Thoroughly test all instructions before deployment

## Development Workflow

1. **Local Development**:
   ```bash
   anchor build
   anchor test
   ```

2. **Testing**:
   ```bash
   anchor test
   ```

3. **Deployment**:
   ```bash
   anchor deploy
   ```

4. **Client Integration**:
   ```javascript
   // Use the generated IDL and program client
   const program = new anchor.Program(idl, programId);
   ```

## Troubleshooting

### Common Issues

1. **Build Errors**: Ensure all dependencies are properly specified in Cargo.toml
2. **Deployment Failures**: Check Solana cluster status and wallet configuration
3. **Client Errors**: Verify program ID and IDL file paths
4. **Test Failures**: Ensure local validator is running and properly configured

### Debug Tips

- Use `msg!()` for logging in Rust programs
- Check transaction logs in Solana Explorer
- Verify account states using `solana account` command
- Use Anchor's built-in debugging tools

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Submit a pull request

## License

This project is licensed under the MIT License.

## Support

For issues and questions:
- Check the [Anchor documentation](https://book.anchor-lang.com/)
- Review [Solana documentation](https://docs.solana.com/)
- Open an issue in the project repository