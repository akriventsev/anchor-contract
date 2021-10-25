use anchor_lang::prelude::*;
declare_id!("Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS");

#[program]
pub mod first_contract {
    use super::*;
    pub fn initialize(ctx: Context<Initialize>) -> ProgramResult {
        msg!("Karamba");
        Ok(())
    }
    pub fn initialize2(ctx: Context<Initialize2>) -> ProgramResult {
        
        //let base_account = Pubkey::new_unique();

        
        
        let decoded = bs58::decode("2eGVwJcrrEWResAeRBB6AXrksgN1WHss9eWUUQgXz8Ux").into_vec().unwrap();
        let addr = &Pubkey::new(decoded.as_slice());
        let owner_decoded = bs58::decode("Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS").into_vec().unwrap();
        let owner = &Pubkey::new(owner_decoded.as_slice());
        let lamports: &mut u64 = &mut 665;
        let data: &mut[u8] = &mut[1, 2, 3];
        let info =  AccountInfo::new(addr,false,false, lamports,
            data,
            owner,
            false,
            56
        );

        let seed = bs58::decode("2eGVwJc").into_vec().unwrap();
        let s: Vec<&[u8]> = vec![seed.as_slice()];
        let seed_slice = s.as_slice();
        let new_addr = Pubkey::create_program_address(seed_slice, &addr)?;

        msg!("Karamba2");
        //msg!(&addr.to_string());
        //base_account.log();
        Ok(())
    }
}

#[derive(Accounts)]
pub struct Initialize{
}

#[derive(Accounts)]
pub struct Initialize2{
}