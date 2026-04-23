#![no_std]
#![no_main]

use core::panic::PanicInfo;

#[panic_handler]
fn panic(_info: &PanicInfo) -> ! {
    loop {}
}

#[unsafe(no_mangle)]
#[unsafe(link_section = "xdp")]
pub extern "C" fn xdp_firewall(_ctx: *mut u8) -> i32 {
    2
}
