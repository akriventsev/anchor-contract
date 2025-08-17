# Usage Examples

## Table of Contents

1. [Basic Setup](#basic-setup)
2. [Solana Program Usage](#solana-program-usage)
3. [JavaScript Client Usage](#javascript-client-usage)
4. [Testing Examples](#testing-examples)
5. [Integration Examples](#integration-examples)
6. [Error Handling Examples](#error-handling-examples)

## Basic Setup

### Prerequisites

1. **Install Solana CLI**
   ```bash
   sh -c "$(curl -sSfL https://release.solana.com/stable/install)"
   ```

2. **Install Anchor CLI**
   ```bash
   npm install -g @project-serum/anchor-cli
   ```

3. **Install Node.js Dependencies**
   ```bash
   npm install
   ```

### Environment Setup

1. **Create Solana Keypair**
   ```bash
   solana-keygen new
   ```

2. **Set Solana Cluster**
   ```bash
   # For local development
   solana config set --url localhost
   
   # For devnet
   solana config set --url devnet
   
   # For mainnet
   solana config set --url mainnet-beta
   ```

3. **Airdrop SOL (for devnet/local)**
   ```bash
   solana airdrop 2
   ```

## Solana Program Usage

### Building the Program

```bash
# Build the program
anchor build

# Build without linting (faster)
anchor build --skip-lint
```

### Deploying the Program

```bash
# Deploy to local cluster
anchor deploy

# Deploy to devnet
anchor deploy --provider.cluster devnet

# Deploy to mainnet
anchor deploy --provider.cluster mainnet
```

### Program Interaction via CLI

```bash
# Call initialize instruction
solana program call <PROGRAM_ID> -- initialize

# Call initialize2 instruction
solana program call <PROGRAM_ID> -- initialize2
```

## JavaScript Client Usage

### Basic Client Setup

```javascript
const anchor = require('@project-serum/anchor');
const { Connection, PublicKey } = require('@solana/web3.js');

// 1. Configure provider
anchor.setProvider(anchor.Provider.env());

// 2. Load IDL
const idl = JSON.parse(require('fs').readFileSync('./target/idl/first_contract.json', 'utf8'));

// 3. Set program ID
const programId = new PublicKey('Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS');

// 4. Create program instance
const program = new anchor.Program(idl, programId);
```

### Calling Program Instructions

#### Example 1: Basic Initialize Call

```javascript
async function callInitialize() {
    try {
        console.log('Calling initialize instruction...');
        
        const tx = await program.rpc.initialize();
        
        console.log('✅ Initialize successful!');
        console.log('Transaction signature:', tx);
        
        return tx;
    } catch (error) {
        console.error('❌ Initialize failed:', error);
        throw error;
    }
}

// Usage
callInitialize()
    .then(tx => console.log('Transaction completed:', tx))
    .catch(err => console.error('Error:', err));
```

#### Example 2: Initialize2 with Error Handling

```javascript
async function callInitialize2() {
    try {
        console.log('Calling initialize2 instruction...');
        
        const tx = await program.rpc.initialize2();
        
        console.log('✅ Initialize2 successful!');
        console.log('Transaction signature:', tx);
        
        return tx;
    } catch (error) {
        console.error('❌ Initialize2 failed:', error);
        
        // Handle specific error types
        if (error.error?.errorCode) {
            console.error('Anchor Error Code:', error.error.errorCode);
        }
        
        if (error.logs) {
            console.error('Program Logs:', error.logs);
        }
        
        throw error;
    }
}

// Usage with retry logic
async function callInitialize2WithRetry(maxRetries = 3) {
    for (let attempt = 1; attempt <= maxRetries; attempt++) {
        try {
            console.log(`Attempt ${attempt}/${maxRetries}`);
            return await callInitialize2();
        } catch (error) {
            if (attempt === maxRetries) {
                throw error;
            }
            console.log(`Attempt ${attempt} failed, retrying...`);
            await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
        }
    }
}
```

### Advanced Client Usage

#### Example 3: Batch Instructions

```javascript
async function batchInstructions() {
    try {
        console.log('Executing batch instructions...');
        
        // Execute multiple instructions in sequence
        const results = [];
        
        // First instruction
        const tx1 = await program.rpc.initialize();
        results.push({ instruction: 'initialize', tx: tx1 });
        console.log('✅ Initialize completed');
        
        // Wait for confirmation
        await program.provider.connection.confirmTransaction(tx1);
        
        // Second instruction
        const tx2 = await program.rpc.initialize2();
        results.push({ instruction: 'initialize2', tx: tx2 });
        console.log('✅ Initialize2 completed');
        
        console.log('🎉 All instructions completed successfully!');
        return results;
        
    } catch (error) {
        console.error('❌ Batch execution failed:', error);
        throw error;
    }
}
```

#### Example 4: Transaction Monitoring

```javascript
async function monitorTransaction(txSignature) {
    try {
        console.log('Monitoring transaction:', txSignature);
        
        // Get transaction details
        const tx = await program.provider.connection.getTransaction(txSignature, {
            commitment: 'confirmed'
        });
        
        if (tx) {
            console.log('✅ Transaction confirmed');
            console.log('Block time:', new Date(tx.blockTime * 1000));
            console.log('Fee:', tx.meta.fee);
            
            // Log program messages
            if (tx.meta.logMessages) {
                console.log('Program logs:');
                tx.meta.logMessages.forEach(log => {
                    if (log.includes('Program log:')) {
                        console.log('  📝', log);
                    }
                });
            }
        } else {
            console.log('⏳ Transaction not yet confirmed');
        }
        
        return tx;
        
    } catch (error) {
        console.error('❌ Error monitoring transaction:', error);
        throw error;
    }
}

// Usage
callInitialize()
    .then(tx => monitorTransaction(tx))
    .catch(err => console.error('Error:', err));
```

## Testing Examples

### Basic Test Structure

```javascript
const anchor = require('@project-serum/anchor');
const { expect } = require('chai');

describe('First Contract Tests', () => {
    let program;
    
    before(async () => {
        // Configure provider
        anchor.setProvider(anchor.Provider.env());
        
        // Get program from workspace
        program = anchor.workspace.FirstContract;
    });
    
    describe('Initialize Instruction', () => {
        it('Should execute initialize successfully', async () => {
            const tx = await program.rpc.initialize();
            expect(tx).to.be.a('string');
            expect(tx.length).to.be.greaterThan(0);
        });
        
        it('Should log "Karamba" message', async () => {
            const tx = await program.rpc.initialize();
            
            // Get transaction details to verify logs
            const txDetails = await program.provider.connection.getTransaction(tx, {
                commitment: 'confirmed'
            });
            
            expect(txDetails.meta.logMessages).to.include('Program log: Karamba');
        });
    });
    
    describe('Initialize2 Instruction', () => {
        it('Should execute initialize2 successfully', async () => {
            const tx = await program.rpc.initialize2();
            expect(tx).to.be.a('string');
            expect(tx.length).to.be.greaterThan(0);
        });
        
        it('Should log "Karamba2" message', async () => {
            const tx = await program.rpc.initialize2();
            
            const txDetails = await program.provider.connection.getTransaction(tx, {
                commitment: 'confirmed'
            });
            
            expect(txDetails.meta.logMessages).to.include('Program log: Karamba2');
        });
    });
});
```

### Advanced Testing

#### Example 5: Test with Custom Assertions

```javascript
describe('Advanced Contract Tests', () => {
    let program;
    
    before(async () => {
        anchor.setProvider(anchor.Provider.env());
        program = anchor.workspace.FirstContract;
    });
    
    it('Should handle multiple rapid calls', async () => {
        const promises = [];
        
        // Make 5 rapid calls
        for (let i = 0; i < 5; i++) {
            promises.push(program.rpc.initialize());
        }
        
        const results = await Promise.all(promises);
        
        expect(results).to.have.length(5);
        results.forEach(tx => {
            expect(tx).to.be.a('string');
            expect(tx.length).to.be.greaterThan(0);
        });
    });
    
    it('Should maintain transaction order', async () => {
        const txs = [];
        
        // Execute instructions sequentially
        for (let i = 0; i < 3; i++) {
            const tx = await program.rpc.initialize();
            txs.push(tx);
            
            // Wait for confirmation
            await program.provider.connection.confirmTransaction(tx);
        }
        
        // Verify all transactions are different
        const uniqueTxs = new Set(txs);
        expect(uniqueTxs.size).to.equal(3);
    });
});
```

## Integration Examples

### Example 6: Web Application Integration

```javascript
// frontend.js - Example frontend integration
class SolanaApp {
    constructor() {
        this.program = null;
        this.connected = false;
    }
    
    async connect() {
        try {
            // Check if Phantom wallet is available
            if (!window.solana || !window.solana.isPhantom) {
                throw new Error('Phantom wallet not found');
            }
            
            // Connect to wallet
            const response = await window.solana.connect();
            this.publicKey = response.publicKey;
            this.connected = true;
            
            console.log('Connected to wallet:', this.publicKey.toString());
            
            // Initialize program
            await this.initializeProgram();
            
        } catch (error) {
            console.error('Connection failed:', error);
            throw error;
        }
    }
    
    async initializeProgram() {
        try {
            // Load IDL and create program
            const idl = await this.loadIDL();
            this.program = new anchor.Program(idl, PROGRAM_ID, this.provider);
            
            console.log('Program initialized');
            
        } catch (error) {
            console.error('Program initialization failed:', error);
            throw error;
        }
    }
    
    async executeInitialize() {
        if (!this.program) {
            throw new Error('Program not initialized');
        }
        
        try {
            const tx = await this.program.rpc.initialize();
            console.log('Initialize executed:', tx);
            return tx;
            
        } catch (error) {
            console.error('Initialize failed:', error);
            throw error;
        }
    }
    
    async executeInitialize2() {
        if (!this.program) {
            throw new Error('Program not initialized');
        }
        
        try {
            const tx = await this.program.rpc.initialize2();
            console.log('Initialize2 executed:', tx);
            return tx;
            
        } catch (error) {
            console.error('Initialize2 failed:', error);
            throw error;
        }
    }
}

// Usage in HTML
/*
<button onclick="app.executeInitialize()">Execute Initialize</button>
<button onclick="app.executeInitialize2()">Execute Initialize2</button>
*/
```

### Example 7: Server-Side Integration

```javascript
// server.js - Example server integration
const express = require('express');
const anchor = require('@project-serum/anchor');

const app = express();
app.use(express.json());

class SolanaService {
    constructor() {
        this.program = null;
        this.initialized = false;
    }
    
    async initialize() {
        try {
            // Configure provider
            anchor.setProvider(anchor.Provider.env());
            
            // Load IDL and create program
            const idl = JSON.parse(require('fs').readFileSync('./target/idl/first_contract.json', 'utf8'));
            const programId = new anchor.web3.PublicKey('Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS');
            
            this.program = new anchor.Program(idl, programId);
            this.initialized = true;
            
            console.log('Solana service initialized');
            
        } catch (error) {
            console.error('Service initialization failed:', error);
            throw error;
        }
    }
    
    async executeInitialize() {
        if (!this.initialized) {
            throw new Error('Service not initialized');
        }
        
        try {
            const tx = await this.program.rpc.initialize();
            return { success: true, transaction: tx };
            
        } catch (error) {
            return { success: false, error: error.message };
        }
    }
    
    async executeInitialize2() {
        if (!this.initialized) {
            throw new Error('Service not initialized');
        }
        
        try {
            const tx = await this.program.rpc.initialize2();
            return { success: true, transaction: tx };
            
        } catch (error) {
            return { success: false, error: error.message };
        }
    }
}

const solanaService = new SolanaService();

// Initialize service on startup
solanaService.initialize().catch(console.error);

// API endpoints
app.post('/api/initialize', async (req, res) => {
    try {
        const result = await solanaService.executeInitialize();
        res.json(result);
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});

app.post('/api/initialize2', async (req, res) => {
    try {
        const result = await solanaService.executeInitialize2();
        res.json(result);
    } catch (error) {
        res.status(500).json({ error: error.message });
    }
});

app.listen(3000, () => {
    console.log('Server running on port 3000');
});
```

## Error Handling Examples

### Example 8: Comprehensive Error Handling

```javascript
class ErrorHandler {
    static handleAnchorError(error) {
        if (error.error?.errorCode) {
            const errorCode = error.error.errorCode;
            
            switch (errorCode.code) {
                case 'InvalidArgument':
                    return {
                        type: 'VALIDATION_ERROR',
                        message: 'Invalid argument provided',
                        details: errorCode.msg
                    };
                    
                case 'InsufficientFunds':
                    return {
                        type: 'INSUFFICIENT_FUNDS',
                        message: 'Insufficient funds for transaction',
                        details: 'Please ensure your wallet has enough SOL'
                    };
                    
                case 'AccountNotInitialized':
                    return {
                        type: 'ACCOUNT_ERROR',
                        message: 'Account not properly initialized',
                        details: 'The account needs to be initialized first'
                    };
                    
                default:
                    return {
                        type: 'ANCHOR_ERROR',
                        message: 'Anchor program error',
                        details: errorCode.msg || 'Unknown error'
                    };
            }
        }
        
        return {
            type: 'UNKNOWN_ERROR',
            message: 'Unknown error occurred',
            details: error.message
        };
    }
    
    static async executeWithErrorHandling(operation, retries = 3) {
        for (let attempt = 1; attempt <= retries; attempt++) {
            try {
                return await operation();
                
            } catch (error) {
                const errorInfo = this.handleAnchorError(error);
                
                console.error(`Attempt ${attempt}/${retries} failed:`, errorInfo);
                
                if (attempt === retries) {
                    throw errorInfo;
                }
                
                // Wait before retry
                await new Promise(resolve => setTimeout(resolve, 1000 * attempt));
            }
        }
    }
}

// Usage
async function safeExecute() {
    try {
        const result = await ErrorHandler.executeWithErrorHandling(
            () => program.rpc.initialize()
        );
        
        console.log('Success:', result);
        return result;
        
    } catch (error) {
        console.error('Operation failed:', error);
        
        // Handle specific error types
        switch (error.type) {
            case 'INSUFFICIENT_FUNDS':
                console.log('Please add more SOL to your wallet');
                break;
                
            case 'VALIDATION_ERROR':
                console.log('Please check your input parameters');
                break;
                
            default:
                console.log('An unexpected error occurred');
        }
        
        throw error;
    }
}
```

### Example 9: Transaction Status Monitoring

```javascript
class TransactionMonitor {
    constructor(connection) {
        this.connection = connection;
    }
    
    async waitForConfirmation(signature, commitment = 'confirmed') {
        try {
            console.log('Waiting for transaction confirmation...');
            
            const confirmation = await this.connection.confirmTransaction(signature, commitment);
            
            if (confirmation.value.err) {
                throw new Error(`Transaction failed: ${confirmation.value.err}`);
            }
            
            console.log('✅ Transaction confirmed');
            return confirmation;
            
        } catch (error) {
            console.error('❌ Transaction confirmation failed:', error);
            throw error;
        }
    }
    
    async getTransactionStatus(signature) {
        try {
            const tx = await this.connection.getTransaction(signature, {
                commitment: 'confirmed'
            });
            
            if (!tx) {
                return { status: 'PENDING', message: 'Transaction not yet confirmed' };
            }
            
            if (tx.meta.err) {
                return { status: 'FAILED', message: 'Transaction failed', error: tx.meta.err };
            }
            
            return { status: 'SUCCESS', message: 'Transaction successful' };
            
        } catch (error) {
            return { status: 'ERROR', message: 'Error checking status', error: error.message };
        }
    }
    
    async monitorTransaction(signature, maxAttempts = 30) {
        for (let attempt = 1; attempt <= maxAttempts; attempt++) {
            const status = await this.getTransactionStatus(signature);
            
            console.log(`Attempt ${attempt}/${maxAttempts}: ${status.status}`);
            
            if (status.status === 'SUCCESS') {
                console.log('🎉 Transaction successful!');
                return status;
            }
            
            if (status.status === 'FAILED') {
                throw new Error(`Transaction failed: ${status.error}`);
            }
            
            // Wait before next check
            await new Promise(resolve => setTimeout(resolve, 2000));
        }
        
        throw new Error('Transaction confirmation timeout');
    }
}

// Usage
const monitor = new TransactionMonitor(program.provider.connection);

callInitialize()
    .then(async (tx) => {
        console.log('Transaction sent:', tx);
        
        try {
            const status = await monitor.monitorTransaction(tx);
            console.log('Final status:', status);
        } catch (error) {
            console.error('Monitoring failed:', error);
        }
    })
    .catch(err => console.error('Error:', err));
```

## Performance Optimization Examples

### Example 10: Batch Processing

```javascript
class BatchProcessor {
    constructor(program) {
        this.program = program;
        this.batchSize = 10;
    }
    
    async processBatch(operations) {
        const results = [];
        const batches = this.chunkArray(operations, this.batchSize);
        
        for (let i = 0; i < batches.length; i++) {
            console.log(`Processing batch ${i + 1}/${batches.length}`);
            
            const batchResults = await Promise.all(
                batches[i].map(op => op())
            );
            
            results.push(...batchResults);
            
            // Small delay between batches to avoid rate limiting
            if (i < batches.length - 1) {
                await new Promise(resolve => setTimeout(resolve, 100));
            }
        }
        
        return results;
    }
    
    chunkArray(array, size) {
        const chunks = [];
        for (let i = 0; i < array.length; i += size) {
            chunks.push(array.slice(i, i + size));
        }
        return chunks;
    }
}

// Usage
const processor = new BatchProcessor(program);

const operations = Array(25).fill(null).map((_, i) => 
    () => program.rpc.initialize()
);

processor.processBatch(operations)
    .then(results => {
        console.log(`Processed ${results.length} operations`);
        console.log('Results:', results);
    })
    .catch(err => console.error('Batch processing failed:', err));
```

These examples demonstrate various ways to use the Solana program and JavaScript client, from basic usage to advanced integration patterns. Each example includes proper error handling and follows best practices for Solana development.