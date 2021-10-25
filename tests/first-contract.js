const anchor = require('@project-serum/anchor');

describe('first-contract', () => {

  // Configure the client to use the local cluster.
  anchor.setProvider(anchor.Provider.env());

  it('Is initialized!', async () => {
    // Add your test here.
    const program = anchor.workspace.FirstContract;
    const tx = await program.rpc.initialize();
    console.log("Your transaction signature", tx);
  });
  it('Is initialized2!', async () => {
    // Add your test here.
    const program = anchor.workspace.FirstContract;
    const tx = await program.rpc.initialize2();
    console.log("Your transaction signature", tx);
  });
});
