// client.js is used to introduce the reader to generating clients from IDLs.
// It is not expected users directly test with this example. For a more
// ergonomic example, see `tests/basic-0.js` in this workspace.

const anchor = require('@project-serum/anchor');

// Configure the local cluster.
anchor.setProvider(anchor.Provider.local('https://api.devnet.solana.com',anchor.Provider.defaultOptions()));

async function main() {
  // #region main
  // Read the generated IDL.
  const idl = JSON.parse(require('fs').readFileSync('./target/idl/first_contract.json', 'utf8'));

  // Address of the deployed program.
  const programId = new anchor.web3.PublicKey('2wJJ6nyY36B1NDB6KS2AQU93YAdoXZgVAm2ShqHv3w69');

  // Generate the program client from IDL.
  const program = new anchor.Program(idl, programId);

  // Execute the RPC.
  await program.rpc.initialize();
  // #endregion main
}

console.log('Running client.');
main().then(() => console.log('Success'));