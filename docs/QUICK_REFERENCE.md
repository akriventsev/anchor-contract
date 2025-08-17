# Quick Reference Guide

## Solana Program Functions

| Function | Description | Parameters | Returns | Context |
|----------|-------------|------------|---------|---------|
| `initialize()` | Logs "Karamba" message | `Context<Initialize>` | `ProgramResult` | `Initialize` |
| `initialize2()` | PDA derivation demo | `Context<Initialize2>` | `ProgramResult` | `Initialize2` |

## Account Structures

| Struct | Fields | Validation |
|--------|--------|------------|
| `Initialize` | None | None required |
| `Initialize2` | None | None required |

## JavaScript Client Methods

| Method | Description | Returns | Example |
|--------|-------------|---------|---------|
| `program.rpc.initialize()` | Execute initialize instruction | `Promise<string>` | `await program.rpc.initialize()` |
| `program.rpc.initialize2()` | Execute initialize2 instruction | `Promise<string>` | `await program.rpc.initialize2()` |

## Provider Configuration

| Method | Description | Example |
|--------|-------------|---------|
| `anchor.setProvider()` | Set global provider | `anchor.setProvider(anchor.Provider.env())` |
| `Provider.local()` | Local cluster provider | `Provider.local('http://localhost:8899')` |
| `Provider.env()` | Environment-based provider | `Provider.env()` |

## Common Commands

### Build & Deploy
```bash
anchor build          # Build program
anchor deploy         # Deploy to local
anchor test           # Run tests
```

### Solana CLI
```bash
solana config set --url localhost    # Set cluster
solana airdrop 2                     # Get test SOL
solana balance                        # Check balance
```

## Error Types

| Error | Description | Handling |
|-------|-------------|----------|
| `ProgramResult::Err` | Program execution error | Check error logs |
| `AnchorError` | Anchor framework error | Use error.errorCode |
| `SolanaError` | Blockchain error | Check transaction status |

## Quick Setup

```javascript
// 1. Configure provider
anchor.setProvider(anchor.Provider.env());

// 2. Load IDL
const idl = JSON.parse(require('fs').readFileSync('./target/idl/first_contract.json', 'utf8'));

// 3. Create program
const program = new anchor.Program(idl, programId);

// 4. Execute instruction
const tx = await program.rpc.initialize();
```

## Testing Pattern

```javascript
describe('Contract Tests', () => {
    let program;
    
    before(async () => {
        anchor.setProvider(anchor.Provider.env());
        program = anchor.workspace.FirstContract;
    });
    
    it('Should work', async () => {
        const tx = await program.rpc.initialize();
        expect(tx).to.be.a('string');
    });
});
```

## File Structure

```
├── programs/first-contract/src/lib.rs    # Main program
├── tests/first-contract.js               # Tests
├── client.js                             # Client example
├── Anchor.toml                           # Anchor config
└── Cargo.toml                            # Rust config
```

## Dependencies

### Rust
- `anchor-lang = "0.17.0"`
- `bs58 = "0.4.0"`

### JavaScript
- `@project-serum/anchor = "^0.17.0"`
- `chai = "^4.3.4"`
- `mocha = "^9.0.3"`

## Program ID
```
Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS
```

## Common Patterns

### Error Handling
```javascript
try {
    const tx = await program.rpc.initialize();
    console.log('Success:', tx);
} catch (error) {
    if (error.error?.errorCode) {
        console.error('Anchor Error:', error.error.errorCode);
    }
}
```

### Transaction Monitoring
```javascript
const tx = await program.rpc.initialize();
await program.provider.connection.confirmTransaction(tx);
```

### Batch Operations
```javascript
const promises = [1, 2, 3].map(() => program.rpc.initialize());
const results = await Promise.all(promises);
```