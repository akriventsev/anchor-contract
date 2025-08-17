# API Reference

## Table of Contents

1. [Solana Program API](#solana-program-api)
2. [JavaScript Client API](#javascript-client-api)
3. [Account Structures](#account-structures)
4. [Error Handling](#error-handling)
5. [Types and Interfaces](#types-and-interfaces)

## Solana Program API

### Program Module: `first_contract`

**Module ID**: `Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS`

#### Function: `initialize`

**Signature**:
```rust
pub fn initialize(ctx: Context<Initialize>) -> ProgramResult
```

**Description**: Executes a simple instruction that logs "Karamba" to the program log.

**Parameters**:
- `ctx: Context<Initialize>` - The instruction context containing account information

**Returns**: `ProgramResult` - Success or error result

**Context Requirements**: `Initialize` struct (no account validation required)

**Example**:
```rust
// Call the initialize function
let result = initialize(ctx)?;
```

**Side Effects**: Logs "Karamba" to the program log

**Error Conditions**: None specific to this function

---

#### Function: `initialize2`

**Signature**:
```rust
pub fn initialize2(ctx: Context<Initialize2>) -> ProgramResult
```

**Description**: Demonstrates program address derivation, base58 decoding, and account manipulation.

**Parameters**:
- `ctx: Context<Initialize2>` - The instruction context containing account information

**Returns**: `ProgramResult` - Success or error result

**Context Requirements**: `Initialize2` struct (no account validation required)

**Functionality**:
1. Decodes a hardcoded base58-encoded public key
2. Creates a program-derived address (PDA) using a seed
3. Demonstrates account info manipulation
4. Logs "Karamba2" to the program log

**Example**:
```rust
// Call the initialize2 function
let result = initialize2(ctx)?;
```

**Side Effects**: 
- Logs "Karamba2" to the program log
- Performs program address derivation

**Error Conditions**:
- Base58 decoding failures
- Program address derivation failures

**Internal Operations**:
```rust
// Decode base58 address
let decoded = bs58::decode("2eGVwJcrrEWResAeRBB6AXrksgN1WHss9eWUUQgXz8Ux").into_vec().unwrap();
let addr = &Pubkey::new(decoded.as_slice());

// Create program-derived address
let seed = bs58::decode("2eGVwJc").into_vec().unwrap();
let s: Vec<&[u8]> = vec![seed.as_slice()];
let seed_slice = s.as_slice();
let new_addr = Pubkey::create_program_address(seed_slice, &addr)?;
```

## JavaScript Client API

### Class: `Program`

**Constructor**:
```javascript
new anchor.Program(idl, programId, provider?)
```

**Parameters**:
- `idl` - The program's Interface Definition Language (IDL)
- `programId` - The program's public key
- `provider` - Optional provider for the connection

**Example**:
```javascript
const program = new anchor.Program(idl, programId);
```

#### Method: `program.rpc.initialize()`

**Signature**:
```javascript
program.rpc.initialize(): Promise<string>
```

**Description**: Executes the `initialize` instruction on the Solana program.

**Parameters**: None

**Returns**: `Promise<string>` - Transaction signature

**Example**:
```javascript
try {
    const tx = await program.rpc.initialize();
    console.log("Transaction signature:", tx);
} catch (error) {
    console.error("Error:", error);
}
```

**Error Handling**: Throws an error if the transaction fails

---

#### Method: `program.rpc.initialize2()`

**Signature**:
```javascript
program.rpc.initialize2(): Promise<string>
```

**Description**: Executes the `initialize2` instruction on the Solana program.

**Parameters**: None

**Returns**: `Promise<string>` - Transaction signature

**Example**:
```javascript
try {
    const tx = await program.rpc.initialize2();
    console.log("Transaction signature:", tx);
} catch (error) {
    console.error("Error:", error);
}
```

**Error Handling**: Throws an error if the transaction fails

### Provider Configuration

#### Function: `anchor.setProvider()`

**Signature**:
```javascript
anchor.setProvider(provider: Provider): void
```

**Description**: Sets the global provider for Anchor operations.

**Parameters**:
- `provider` - The provider instance to use

**Example**:
```javascript
// For local cluster
anchor.setProvider(anchor.Provider.local('https://api.devnet.solana.com', anchor.Provider.defaultOptions()));

// For environment (recommended for testing)
anchor.setProvider(anchor.Provider.env());
```

#### Class: `anchor.Provider`

**Static Method**: `Provider.local()`

**Signature**:
```javascript
Provider.local(endpoint: string, options?: ProviderOptions): Provider
```

**Description**: Creates a provider for a local Solana cluster.

**Parameters**:
- `endpoint` - The RPC endpoint URL
- `options` - Optional provider configuration

**Returns**: `Provider` instance

**Example**:
```javascript
const provider = anchor.Provider.local('https://api.devnet.solana.com');
```

**Static Method**: `Provider.env()`

**Signature**:
```javascript
Provider.env(): Provider
```

**Description**: Creates a provider using environment variables.

**Returns**: `Provider` instance

**Example**:
```javascript
const provider = anchor.Provider.env();
```

## Account Structures

### Struct: `Initialize`

**Definition**:
```rust
#[derive(Accounts)]
pub struct Initialize {
    // No account validation required
}
```

**Description**: Account validation structure for the `initialize` instruction.

**Fields**: None

**Validation**: No account validation required

**Usage**:
```rust
pub fn initialize(ctx: Context<Initialize>) -> ProgramResult {
    // Function implementation
}
```

---

### Struct: `Initialize2`

**Definition**:
```rust
#[derive(Accounts)]
pub struct Initialize2 {
    // No account validation required
}
```

**Description**: Account validation structure for the `initialize2` instruction.

**Fields**: None

**Validation**: No account validation required

**Usage**:
```rust
pub fn initialize2(ctx: Context<Initialize2>) -> ProgramResult {
    // Function implementation
}
```

## Error Handling

### Error Types

#### `ProgramResult`
**Type**: `Result<(), ProgramError>`

**Description**: The return type for Solana program functions.

**Values**:
- `Ok(())` - Function executed successfully
- `Err(ProgramError)` - Function failed with an error

**Example**:
```rust
pub fn my_function() -> ProgramResult {
    if some_condition {
        Ok(())
    } else {
        Err(ProgramError::InvalidArgument)
    }
}
```

#### `ProgramError`
**Description**: Enumeration of possible program errors.

**Common Variants**:
- `InvalidArgument` - Invalid argument provided
- `InsufficientFunds` - Insufficient funds for operation
- `AccountNotInitialized` - Account not properly initialized
- `Custom(ErrorCode)` - Custom error with specific code

### Error Handling Patterns

#### Rust Error Handling
```rust
pub fn safe_function() -> ProgramResult {
    // Perform operation that might fail
    let result = risky_operation()?;
    
    // Continue with success
    Ok(())
}
```

#### JavaScript Error Handling
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

## Types and Interfaces

### Solana Types

#### `Pubkey`
**Description**: Represents a public key in the Solana system.

**Creation**:
```rust
// From bytes
let pubkey = Pubkey::new(&[1, 2, 3, 4]);

// From base58 string
let decoded = bs58::decode("base58string").into_vec().unwrap();
let pubkey = Pubkey::new(&decoded);
```

**Methods**:
- `to_string()` - Convert to base58 string
- `new_unique()` - Generate a new unique public key

#### `AccountInfo`
**Description**: Contains information about a Solana account.

**Fields**:
- `key` - Public key of the account
- `is_signer` - Whether the account is a signer
- `is_writable` - Whether the account is writable
- `lamports` - Number of lamports in the account
- `data` - Account data
- `owner` - Program that owns the account
- `executable` - Whether the account is executable
- `rent_epoch` - Epoch when rent was last collected

**Example**:
```rust
let account_info = AccountInfo::new(
    &pubkey,
    false,  // is_signer
    false,  // is_writable
    &mut lamports,
    &mut data,
    &owner,
    false,  // executable
    rent_epoch
);
```

### Anchor Types

#### `Context<T>`
**Description**: Generic context wrapper for instruction parameters.

**Generic Parameter**: `T` - The account validation structure

**Usage**:
```rust
pub fn my_function(ctx: Context<MyAccounts>) -> ProgramResult {
    // Access accounts through ctx.accounts
    let account = &ctx.accounts.my_account;
    Ok(())
}
```

#### `msg!()`
**Macro**: Logs a message to the program log.

**Usage**:
```rust
msg!("Hello, World!");
msg!("Value: {}", 42);
```

**Output**: Appears in transaction logs and program logs